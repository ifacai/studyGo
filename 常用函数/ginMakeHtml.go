func wordIndex(c *gin.Context) {
	htmlExist, htmlPath, htmlFileName := checkHtmlExist(c.Request.RequestURI)
	if htmlExist {
		//html已经存在 直接返回
		htmlCode := readFileString(htmlPath + htmlFileName)
		c.Writer.WriteHeader(http.StatusOK)
		if _, err := c.Writer.WriteString(htmlCode); err != nil {
			panic(err)
		}
	} else {
		//html不存在 开始生成
		parent, child := getAllCateByMysql()
		c.HTML(http.StatusOK, "wordIndex.html", gin.H{
			"parent": parent,
			"child":  child,
		})
	
		myTemp, _ := template.ParseFiles("./template/wordIndex.html", "./template/header.html", "./template/footer.html")
		makeHtml(myTemp, htmlPath, htmlFileName, gin.H{
			"parent": parent,
			"child":  child,
		})
	}
}

func checkHtmlExist(requestUrl string) (exist bool, htmlPath, htmlFileName string) {
	md5String := md5Encode(requestUrl)
	_rune := []rune(md5String)
	firstDir := string(_rune[0:1])
	secondDir := string(_rune[1:3])
	htmlPath = "/data/html/" + firstDir + "/" + secondDir + "/"
	htmlFileName = md5String + ".html"
	exist = checkFileExist(htmlPath + htmlFileName)
	return exist, htmlPath, htmlFileName
}

func readFileString(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("error: ", err)
		return ""
	}
	//defer f.Close()
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	fd, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}
	return string(fd)
}

func makeHtml(myTemplate *template.Template, path, fileName string, params gin.H) {
	// path like "./html/market/"
	var f *os.File
	fullFileName := path + fileName
	pathExist := checkFileExist(path)
	if !pathExist {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.Create(fullFileName)
	if err != nil {
		panic(fullFileName + " 创建失败")
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	err = myTemplate.Execute(f, params)
	if err != nil {
		fmt.Println(err)
		panic("")
	}
}
