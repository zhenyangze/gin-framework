package helpers

import (
	"regexp"
	"strings"
)

func FilterToLower(html string) string {
	reg, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = reg.ReplaceAllStringFunc(html, strings.ToLower)
	return html
}

func FilterHTML(html string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	html = re.ReplaceAllString(html, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	html = re.ReplaceAllString(html, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllString(html, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	html = re.ReplaceAllString(html, "\n")

	return strings.TrimSpace(html)
}

// FilterIframe 过滤iframe
func FilterIframe(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<iframe[\\S\\s]+?\\</iframe\\>")
	html = reg.ReplaceAllString(html, "<p class='style'></p>")
	return html
}

// FilterXML 过滤xml
func FilterXML(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<?xml[\\S\\s]+?\\?\\>")
	html = reg.ReplaceAllString(html, " ")
	return html
}

// FilterStyle 过滤html中的style
func FilterStyle(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	html = reg.ReplaceAllString(html, " ")
	return html
}

// FilterJS 过滤html中的js
func FilterJS(html string) string {
	html = FilterToLower(html)
	reg, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	html = reg.ReplaceAllString(html, " ")
	return html
}
