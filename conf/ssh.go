package conf

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// 创建SSH连接
func (m *Master) CreateSSHConn() (*ssh.Client, error) {
	config := ssh.ClientConfig{
		User: m.SysUsername,
		Auth: []ssh.AuthMethod{ssh.Password(m.SysPassword)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 60 * time.Minute,
	}
	addr := fmt.Sprintf("%s:%d", m.SysHost, m.SysPort)
	client, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 创建sftp连接
func (m *Master) CreateSftpConn() (*sftp.Client, error) {
	sshClient, err := m.CreateSSHConn()
	if err != nil {
		return nil, err
	}
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}
	return sftpClient, nil
}

// 创建SSH会话用于执行shell命令
func (m *Master) RunShell(shell string) (string, error) {
	client, err := m.CreateSSHConn()
	if err != nil {
		return "", err
	}
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	output, err := session.CombinedOutput(shell)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// 创建上传文件的方法
func (m *Master) UploadFile(srcFile, dstFile string) (string, error) {
	sftpClient, err := m.CreateSftpConn()
	if err != nil {
		return "创建sftp连接失败!", err
	}
	// 打开本地文件
	dst, err := sftpClient.Create(dstFile)
	if err != nil {
		return fmt.Sprintf("打开远程文件[%s]失败! 原因: %s", dstFile, err.Error()), err
	}
	defer dst.Close()
	src, err := os.Open(srcFile)
	if err != nil {
		return fmt.Sprintf("打开本地文件[%s]失败! 原因: %s", srcFile, err.Error()), err
	}
	defer src.Close()
	n, err := io.Copy(dst, src)
	if err != nil {
		return fmt.Sprintf("上传本地文件[%s]失败! 原因: %s", srcFile, err.Error()), err
	}
	dstFileInfo, err := sftpClient.Stat(dstFile)
	if err != nil {
		return fmt.Sprintf("获取远程文件[%s]信息失败! 原因: %s", dstFile, err.Error()), err
	}
	if n != dstFileInfo.Size() {
		return fmt.Sprintf("上传本地文件[%s]失败!", srcFile), fmt.Errorf("本地文件和远程文件大小不一样")
	}
	return fmt.Sprintf("上传本地文件[%s]成功!", srcFile), err
}

// 创建下载osw文件的方法
func (s *Master) DownloadFile(srcFile, dstFile string) (string, error) {
	sftpClient, err := s.CreateSftpConn()
	if err != nil {
		return "创建sftp连接失败!", err
	}
	src, err := sftpClient.Open(srcFile)
	if err != nil {
		return fmt.Sprintf("打开远程文件[%s]流失败!", srcFile), err
	}
	dst, err := os.Create(dstFile)
	if err != nil {
		return fmt.Sprintf("打开本地文件[%s]流失败!", dstFile), err
	}
	_, err = src.WriteTo(dst)
	if err != nil {
		return fmt.Sprintf("远程文件[%s]下载失败!", srcFile), err
	}
	return fmt.Sprintf("远程文件[%s]下载成功!", srcFile), err
}

// 创建SSH连接
func (m *Slavea) CreateSSHConn() (*ssh.Client, error) {
	config := ssh.ClientConfig{
		User: m.SysUsername,
		Auth: []ssh.AuthMethod{ssh.Password(m.SysPassword)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 60 * time.Minute,
	}
	addr := fmt.Sprintf("%s:%d", m.SysHost, m.SysPort)
	client, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 创建sftp连接
func (m *Slavea) CreateSftpConn() (*sftp.Client, error) {
	sshClient, err := m.CreateSSHConn()
	if err != nil {
		return nil, err
	}
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}
	return sftpClient, nil
}

// 创建SSH会话用于执行shell命令
func (m *Slavea) RunShell(shell string) (string, error) {
	client, err := m.CreateSSHConn()
	if err != nil {
		return "", err
	}
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	output, err := session.CombinedOutput(shell)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// 创建上传文件的方法
func (m *Slavea) UploadFile(srcFile, dstFile string) (string, error) {
	sftpClient, err := m.CreateSftpConn()
	if err != nil {
		return "创建sftp连接失败!", err
	}
	// 打开本地文件
	dst, err := sftpClient.Create(dstFile)
	if err != nil {
		return fmt.Sprintf("打开远程文件[%s]失败! 原因: %s", dstFile, err.Error()), err
	}
	defer dst.Close()
	src, err := os.Open(srcFile)
	if err != nil {
		return fmt.Sprintf("打开本地文件[%s]失败! 原因: %s", srcFile, err.Error()), err
	}
	defer src.Close()
	n, err := io.Copy(dst, src)
	if err != nil {
		return fmt.Sprintf("上传本地文件[%s]失败! 原因: %s", srcFile, err.Error()), err
	}
	dstFileInfo, err := sftpClient.Stat(dstFile)
	if err != nil {
		return fmt.Sprintf("获取远程文件[%s]信息失败! 原因: %s", dstFile, err.Error()), err
	}
	if n != dstFileInfo.Size() {
		return fmt.Sprintf("上传本地文件[%s]失败!", srcFile), fmt.Errorf("本地文件和远程文件大小不一样")
	}
	return fmt.Sprintf("上传本地文件[%s]成功!", srcFile), err
}

// 创建SSH连接
func (m *Slaveb) CreateSSHConn() (*ssh.Client, error) {
	config := ssh.ClientConfig{
		User: m.SysUsername,
		Auth: []ssh.AuthMethod{ssh.Password(m.SysPassword)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 60 * time.Minute,
	}
	addr := fmt.Sprintf("%s:%d", m.SysHost, m.SysPort)
	client, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 创建sftp连接
func (m *Slaveb) CreateSftpConn() (*sftp.Client, error) {
	sshClient, err := m.CreateSSHConn()
	if err != nil {
		return nil, err
	}
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}
	return sftpClient, nil
}

// 创建SSH会话用于执行shell命令
func (m *Slaveb) RunShell(shell string) (string, error) {
	client, err := m.CreateSSHConn()
	if err != nil {
		return "", err
	}
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	output, err := session.CombinedOutput(shell)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// 创建上传文件的方法
func (m *Slaveb) UploadFile(srcFile, dstFile string) (string, error) {
	sftpClient, err := m.CreateSftpConn()
	if err != nil {
		return "创建sftp连接失败!", err
	}
	// 打开本地文件
	dst, err := sftpClient.Create(dstFile)
	if err != nil {
		return fmt.Sprintf("打开远程文件[%s]失败! 原因: %s", dstFile, err.Error()), err
	}
	defer dst.Close()
	src, err := os.Open(srcFile)
	if err != nil {
		return fmt.Sprintf("打开本地文件[%s]失败! 原因: %s", srcFile, err.Error()), err
	}
	defer src.Close()
	n, err := io.Copy(dst, src)
	if err != nil {
		return fmt.Sprintf("上传本地文件[%s]失败! 原因: %s", srcFile, err.Error()), err
	}
	dstFileInfo, err := sftpClient.Stat(dstFile)
	if err != nil {
		return fmt.Sprintf("获取远程文件[%s]信息失败! 原因: %s", dstFile, err.Error()), err
	}
	if n != dstFileInfo.Size() {
		return fmt.Sprintf("上传本地文件[%s]失败!", srcFile), fmt.Errorf("本地文件和远程文件大小不一样")
	}
	return fmt.Sprintf("上传本地文件[%s]成功!", srcFile), err
}
