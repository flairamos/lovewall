package filebrowser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shoollove/config"
	"strings"
)

func checkFileType(name string) string {
	switch name {
	case "png":
		return "image/png"
	case "jpg", "jpeg":
		return "image/jpg"
	case "webp":
		return "image/webp"
	case "gif":
		return "image/gif"
	default:
		return ""
	}
}

type FileBrowserConfig struct {
	BaseURL  string // File Browser 服务器地址   结尾不需要 /
	Username string // 用户名
	Password string // 密码
	FilePath string // 目标路径   以/ 开头绝对路径相当于 /
	Override bool   // 是否覆盖
}

func NewConfig(url, username, password, filepath string) FileBrowserConfig {
	return FileBrowserConfig{
		BaseURL:  url,
		Username: username,
		Password: password,
		FilePath: filepath,
	}
}

// 获取认证令牌
func GetToken(config FileBrowserConfig) (string, error) {
	// 创建登录请求
	loginData := fmt.Sprintf(`{"username":"%s","password":"%s"}`, config.Username, config.Password)

	// 创建请求
	req, err := http.NewRequest("POST", config.BaseURL+"/api/login", bytes.NewBufferString(loginData))
	if err != nil {
		return "", fmt.Errorf("创建登录请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("发送登录请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应内容失败: %v", err)
	}

	// 打印原始响应内容
	// fmt.Printf("原始响应内容: %s\n", string(bodyBytes))
	return string(bodyBytes), nil
}

func UploadFileStream(config FileBrowserConfig, by []byte) (string, error) {
	reader := bytes.NewReader(by)
	// 确保目标路径以/开头
	var targetPath = config.FilePath
	if !strings.HasPrefix(targetPath, "/") {
		targetPath = "/" + targetPath
	}

	// 构建请求URL
	url := fmt.Sprintf("%s/api/resources%s?override=%v", config.BaseURL, targetPath, config.Override)

	// 取文件名
	fileExt := ""
	split := strings.Split(strings.ToLower(config.FilePath), "/")
	for i := len(split) - 1; i >= 0; i-- {
		if strings.Contains(split[i], ".") {
			str := strings.Split(split[i], ".")
			fileExt = strings.TrimSpace(str[1])
			break
		}
	}
	// 创建请求
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	if checkFileType(fileExt) == "" {
		return "", fmt.Errorf("不支持该文件上传: %v", err)
	}
	// 设置请求头
	req.Header.Set("Content-Type", "image/png")
	token, err := GetToken(config)
	if err != nil {
		return "", fmt.Errorf("Token请求失败: %v", err)
	}
	req.Header.Set("X-Auth", token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("上传失败: %s (状态码: %d)", string(respBody), resp.StatusCode)
	}

	return string(respBody), nil
}

// DeleteFile 删除 File Browser 中的文件
func DeleteFile(filePath string) error {
	// 创建请求体
	requestBody := map[string]string{
		"path": filePath,
	}

	// 将请求体转换为 JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("创建请求体失败: %v", err)
	}

	// 创建请求
	req, err := http.NewRequest("DELETE", config.Config.Filebrowser.Url+"/api/resources/"+config.Config.Filebrowser.Bucket+"/"+filePath, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	fconfig := FileBrowserConfig{
		BaseURL:  config.Config.Filebrowser.Url,
		Username: config.Config.Filebrowser.Username,
		Password: config.Config.Filebrowser.Password,
	}
	token, err := GetToken(fconfig)
	if err != nil {
		return err
	}
	req.Header.Set("x-auth", token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("删除失败，状态码: %d，响应: %s", resp.StatusCode, string(bodyBytes))
	}
	return nil
}
