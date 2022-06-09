package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	months := []string{"farvardin", "ordibehesht", "khordad", "tir", "mordad", "shahrivar", "mehr", "aban", "azar", "dey", "bahman", "esfand"}

	os.Mkdir("c-sharp\\", os.ModeDir)
	os.Mkdir("c-sharp\\leave\\", os.ModeDir)
	os.Mkdir("c-sharp\\payment\\", os.ModeDir)
	me := widget.NewLabel("there isn't any standard data on c-sharp folder do you like example insert for test ?")
	he := widget.NewFormItem("Attention", me)
	leavelistmem := []string{}
	leavelistmemcount := []string{}
	var temp string = "مدیرﻳﺖ" + "کارمندان"
	validletters := "abcdefghijklmnopqrstunvwxyzABSDEFGHIJKLMNOPQRSTUNVWXYNZ. "
	validnumbers := "1234567890"
	a := app.NewWithID(temp)
	w := a.NewWindow(temp)
	var userList []*widget.FormItem = []*widget.FormItem{he}
	d := dialog.NewForm("example datas", "insert example data", "don't insert example data", userList, func(tempbool bool) {
		if tempbool {
			var file string = "c-sharp\\theemployee.txt"
			var tempstr1 string = "1232|Ali|Akbari|09121234567\n11232354534|moslem|moslemi|0999992399\n"
			var fi string = "c-sharp\\"
		tmgotocr6:
			file1, err6 := os.Create(file)
			if err6 != nil {
				goto tmgotocr6
			}
			file1.WriteString(tempstr1)

			file1.Close()

		tmgotocr7:
			file2, err7 := os.Create(fi + "payment\\mordad.txt")
			if err7 != nil {
				goto tmgotocr7
			}
			file2.WriteString("1232|7000000")

			file2.Close()

		tmgotocr8:
			file3, err8 := os.Create(fi + "leave\\mordad.txt")
			if err8 != nil {
				goto tmgotocr8
			}
			file3.WriteString("1232|2")

			file3.Close()

		}

	}, w)
	_, te1 := os.Open("c-sharp\\theemployee.txt")
	if te1 != nil {
		d.Show()
	}
	w2 := a.NewWindow("sds")
	insertLeave := dialog.NewCustom("Leave Days", "close", w2.Content(), w)
	text1 := widget.NewEntry()
	text1.SetPlaceHolder("ID")
	text8 := widget.NewEntry()
	text8.SetPlaceHolder("Name")
	text9 := widget.NewEntry()
	text9.SetPlaceHolder("Family")
	text10 := widget.NewEntry()
	text10.SetPlaceHolder("Phone number")
	text11 := widget.NewLabel("ID : ")

	text2 := widget.NewEntry()
	text5 := widget.NewEntry()
	text6 := widget.NewEntry()
	listleave1 := widget.NewSelect(leavelistmem, func(s string) {
		tempint1 := 0
		for i := 0; i < len(leavelistmem); i++ {
			if leavelistmem[i] == s {
				tempint1 = i
				break
			}
		}
		text5.SetText(leavelistmemcount[tempint1])
	})
	listpay1 := widget.NewSelect(leavelistmem, func(s string) {
		tempint1 := 0
		for i := 0; i < len(leavelistmem); i++ {
			if leavelistmem[i] == s {
				tempint1 = i
				break
			}
		}
		text6.SetText(leavelistmemcount[tempint1])
	})
	text5.SetPlaceHolder("Days of leave")
	text6.SetPlaceHolder("Amount of pay")
	text12 := widget.NewLabel("Family : ")
	text2.SetPlaceHolder("Family")

	text3 := widget.NewEntry()
	text13 := widget.NewLabel("Name : ")
	text3.SetPlaceHolder("Name")
	text4 := widget.NewEntry()
	text14 := widget.NewLabel("Phone number : ")
	text4.SetPlaceHolder("Phone number")

	templist1 := widget.NewSelect(months, func(s string) {
		file := "c-sharp\\leave\\" + s + ".txt"
		leavelistmem = nil
		leavelistmemcount = nil
		f, err := os.Open(file)
		if err != nil {
			listleave1 = widget.NewSelect(leavelistmem, func(s string) {
				tempint1 := 0
				for i := 0; i < len(leavelistmem); i++ {
					if leavelistmem[i] == s {
						tempint1 = i
						break
					}
				}
				text5.SetText(leavelistmemcount[tempint1])
			})
			return
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var tempstr2 string = (scanner.Text())
			for temp1 := 0; temp1 < len(tempstr2); temp1++ {
				if tempstr2[temp1] == '|' {
					tempstr3 := scanner.Text()[:temp1]
					leavelistmem = append(leavelistmem, tempstr3)
					tempstr3 = scanner.Text()[temp1+1:]
					leavelistmemcount = append(leavelistmemcount, tempstr3)
				}
			}

		}

		listleave1 = widget.NewSelect(leavelistmem, func(s string) {
			tempint1 := 0
			for i := 0; i < len(leavelistmem); i++ {
				if leavelistmem[i] == s {
					tempint1 = i
					break
				}
			}
			text5.SetText(leavelistmemcount[tempint1])
		})
		f.Close()
	})
	templist2 := widget.NewSelect(months, func(s string) {
		file := "c-sharp\\payment\\" + s + ".txt"
		leavelistmem = nil
		leavelistmemcount = nil
		f, err := os.Open(file)
		if err != nil {
			listpay1 = widget.NewSelect(leavelistmem, func(s string) {
				tempint1 := 0
				for i := 0; i < len(leavelistmem); i++ {
					if leavelistmem[i] == s {
						tempint1 = i
						break
					}
				}
				text6.SetText(leavelistmemcount[tempint1])
			})
			return
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var tempstr2 string = (scanner.Text())
			for temp1 := 0; temp1 < len(tempstr2); temp1++ {
				if tempstr2[temp1] == '|' {
					tempstr3 := scanner.Text()[:temp1]
					leavelistmem = append(leavelistmem, tempstr3)
					tempstr3 = scanner.Text()[temp1+1:]
					leavelistmemcount = append(leavelistmemcount, tempstr3)
				}
			}

		}

		listpay1 = widget.NewSelect(leavelistmem, func(s string) {
			tempint1 := 0
			for i := 0; i < len(leavelistmem); i++ {
				if leavelistmem[i] == s {
					tempint1 = i
					break
				}
			}
			text6.SetText(leavelistmemcount[tempint1])
		})
		f.Close()
	})

	insertemp := dialog.NewCustom("New Employee", "close", w2.Content(), w)

	tempbut1 := widget.NewButton("Submit", func() {
		tempint1 := 1
		tempint2 := 1
		tempint3 := 1
		tempint5 := 1
		strtemp1 := ""
		tempint4 := 1
		tempfulint1 := 4

		if text1.Text == "" {
			strtemp1 += "Name field must not be empty\n"
			tempfulint1--
			tempint1--
		}
		if text2.Text == "" {
			strtemp1 += "Family field must not be empty\n"
			tempfulint1--
			tempint2--
		}
		if text3.Text == "" {
			strtemp1 += "ID field must not be empty\n"
			tempfulint1--
			tempint3--
		}
		if text4.Text == "" {
			strtemp1 += "Phone number field must not be empty\n"
			tempfulint1--
			tempint4--
		}
		for i := 0; i < len(text1.Text); i++ {
			for j := 0; j < len(validnumbers); j++ {
				if text1.Text[i] == validnumbers[j] {
					break
				}
				if j == len(validnumbers)-1 {
					strtemp1 += "Name field must not have " + string(text1.Text[i]) + " \n"
					tempint1--
					tempfulint1--
				}
			}
		}
		for i := 0; i < len(text2.Text); i++ {
			for j := 0; j < len(validletters); j++ {
				if text2.Text[i] == validletters[j] {
					break
				}
				if j == len(validletters)-1 {
					strtemp1 += "Family field must not have " + string(text2.Text[i]) + " \n"
					tempint2--
					tempfulint1--
				}
			}
		}
		for i := 0; i < len(text3.Text); i++ {
			for j := 0; j < len(validletters); j++ {
				if text3.Text[i] == validletters[j] {
					break
				}
				if j == len(validletters)-1 {
					strtemp1 += "ID field must not have " + string(text3.Text[i]) + " \n"
					tempint3--
					tempfulint1--
				}
			}
		}
		for i := 0; i < len(text4.Text); i++ {
			for j := 0; j < len(validnumbers); j++ {
				if text4.Text[i] == validnumbers[j] {
					break
				}
				if j == len(validnumbers)-1 {
					strtemp1 += "Phone number field must not have " + string(text4.Text[i]) + " \n"
					tempint4--
					tempfulint1--
				}
			}
		}
		checkerrs, _ := os.OpenFile("c-sharp\\theemployee.txt", os.O_RDWR|os.O_CREATE, 0644)
		defer checkerrs.Close()
		scanner := bufio.NewScanner(checkerrs)
		var mainstr1 []string
		for scanner.Scan() {
			var tempstr2 string = (scanner.Text())
			for temp1 := 0; temp1 < len(tempstr2); temp1++ {
				if tempstr2[temp1] == '|' {
					mainstr1 = append(mainstr1, string(scanner.Text()[:temp1]))

					break
				}
			}
		}

		for i := 0; i < len(mainstr1); i++ {
			if mainstr1[i] == string(text1.Text) {
				tempint5--
				break
			}
		}
		if tempint5 != 1 {
			tempdiag1 := dialog.NewError(errors.New("Must not have same ID for insert employee"), w)

			tempdiag1.Show()

			text1.FocusGained()
		} else if tempfulint1 != 4 {
			tempdiag1 := dialog.NewError(errors.New(strtemp1), w)

			tempdiag1.Show()

			if tempint1 != 1 {
				text1.FocusGained()
			}
			if tempint2 != 1 {
				text2.FocusGained()
			}
			if tempint3 != 1 {
				text3.FocusGained()
			}
			if tempint4 != 1 {
				text4.FocusGained()
			}
		} else {

			strtemp1 = "\n" + string(text1.Text) + "|" + string(text2.Text) + "|" + string(text3.Text) + "|" + text4.Text
			f, _ := os.OpenFile("c-sharp\\theemployee.txt",
				os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

			f.WriteString(strtemp1)
			app.NewWithID(temp).SendNotification(&fyne.Notification{
				Title:   temp,
				Content: "Eployee Insert " + text1.Text + "\nName : " + text2.Text + "Family :" + text3.Text + "\n Phone number" + text4.Text,
			})
			text1.SetText("")
			text2.SetText("")
			text3.SetText("")
			text4.SetText("")
			f.Close()
			insertdepay()
			insertdelea()
		}

	})
	insertLeave1 := dialog.NewCustom("", "close", w2.Content(), w)
	insertpay1 := dialog.NewCustom("", "close", w2.Content(), w)
	tempbut2 := widget.NewButton("Clear all", func() { text1.SetText(""); text2.SetText(""); text3.SetText(""); text4.SetText("") })

	editbutleave := widget.NewButton("Edit", func() {
		strtemp1 := ""
		tempint1 := 1
		for i := 0; i < len(text5.Text); i++ {
			for j := 0; j < len(validnumbers); j++ {
				if text5.Text[i] == validnumbers[j] {
					break
				}
				if j == len(validnumbers)-1 {
					strtemp1 += "Leave field must not have " + string(text5.Text[i]) + " \n"
					tempint1--

				}
			}
		}
		if tempint1 != 1 {
			tempdiag1 := dialog.NewError(errors.New(strtemp1), w)

			tempdiag1.Show()

			text5.FocusGained()
		} else {
			var mfile string = "c-sharp\\leave\\" + templist1.Selected + ".txt"
		tmgotoer2:
			f, err := os.Open(mfile)
			if err != nil {
				goto tmgotoer2
			}
			scanner := bufio.NewScanner(f)
			var mainstr1 []string
			var mainstr11 []string
			for scanner.Scan() {
				var tempstr2 string = (scanner.Text())
				for temp1 := 0; temp1 < len(tempstr2); temp1++ {
					if tempstr2[temp1] == '|' {
						mainstr1 = append(mainstr1, string(scanner.Text()[:temp1]))
						mainstr11 = append(mainstr11, string(scanner.Text()[temp1+1:]))
						break
					}
				}
			}

			tempint1 := 0
			for i := 0; i < len(mainstr1); i++ {
				if mainstr1[i] == listleave1.Selected {
					tempint1 = i

					break
				}

			}
			mainstr11[tempint1] = text5.Text
			leavelistmemcount[tempint1] = text5.Text
			strtemp1 := ""
		tmgotocr1:
			s1, err := os.Create(mfile)
			if err != nil {
				goto tmgotocr1
			}
			for temp1 := 0; temp1 < len(mainstr1); temp1++ {
				strtemp1 = strtemp1 + mainstr1[temp1] + "|" + mainstr11[temp1] + "\n"

			}
			s1.WriteString(strtemp1)
			s1.Close()
			f.Close()
		}
	})
	editbutpay := widget.NewButton("Edit", func() {
		strtemp1 := ""
		tempint1 := 1
		for i := 0; i < len(text6.Text); i++ {
			for j := 0; j < len(validnumbers); j++ {
				if text6.Text[i] == validnumbers[j] {
					break
				}
				if j == len(validnumbers)-1 {
					strtemp1 += "Pay field must not have " + string(text6.Text[i]) + " \n"
					tempint1--

				}
			}
		}
		if tempint1 != 1 {
			tempdiag1 := dialog.NewError(errors.New(strtemp1), w)

			tempdiag1.Show()

			text6.FocusGained()
		} else {
			var mfile string = "c-sharp\\payment\\" + templist2.Selected + ".txt"
		tmgotoer2:
			f, err := os.Open(mfile)
			if err != nil {
				goto tmgotoer2
			}

			scanner := bufio.NewScanner(f)
			var mainstr1 []string
			var mainstr11 []string
			for scanner.Scan() {
				var tempstr2 string = (scanner.Text())
				for temp1 := 0; temp1 < len(tempstr2); temp1++ {
					if tempstr2[temp1] == '|' {
						mainstr1 = append(mainstr1, string(scanner.Text()[:temp1]))
						mainstr11 = append(mainstr11, string(scanner.Text()[temp1+1:]))
						break
					}
				}
			}

			tempint1 := 0
			for i := 0; i < len(mainstr1); i++ {
				if mainstr1[i] == listpay1.Selected {
					tempint1 = i

					break
				}

			}
			fmt.Println(tempint1)
			mainstr11[tempint1] = text6.Text
			leavelistmemcount[tempint1] = text6.Text
			strtemp1 := ""
		tmgotocr1:
			s1, err := os.Create(mfile)
			if err != nil {
				goto tmgotocr1
			}
			for temp1 := 0; temp1 < len(mainstr1); temp1++ {
				strtemp1 = strtemp1 + mainstr1[temp1] + "|" + mainstr11[temp1] + "\n"

			}

			s1.WriteString(strtemp1)
			s1.Close()
			f.Close()
		}
	})
	w.SetContent(container.NewAdaptiveGrid(2,

		widget.NewButton("Insert new employee", func() {
			w2.SetContent(container.New(layout.NewFormLayout(), text11, text1, text12, text2, text13, text3, text14, text4, tempbut2, tempbut1))
			insertemp = dialog.NewCustom("New Employee", "close", w2.Content(), w)
			insertemp.Show()

		}), widget.NewButton("Leave days control", func() {
			w2.SetContent(container.NewVBox(templist1, widget.NewButton("show", func() {
				text5.SetText("")
				listleave1.ClearSelected()
				text5.SetText("")
				w2.SetContent(container.NewVBox(listleave1, text5, editbutleave))
				insertLeave1 = dialog.NewCustom(templist1.Selected+" employee leaves", "close", w2.Content(), w)
				insertLeave1.Show()

			})))
			insertLeave = dialog.NewCustom("Leave of employees", "close", w2.Content(), w)
			insertLeave.Show()

		}), widget.NewButton("Payment control", func() {
			w2.SetContent(container.NewVBox(templist2, widget.NewButton("show", func() {
				text6.SetText("")
				listpay1.ClearSelected()
				text6.SetText("")
				w2.SetContent(container.NewVBox(listpay1, text6, editbutpay))
				insertpay1 = dialog.NewCustom(templist2.Selected+" employee payments", "close", w2.Content(), w)
				insertpay1.Show()

			})))
			insertpay1 = dialog.NewCustom("Payment of employees", "close", w2.Content(), w)
			insertpay1.Show()

		}), widget.NewButton("Exit", func() {
			w.Hide()
			w.Close()
			w.Hide()
			w.Close()
		}),
	))
	fyne.Window.Clipboard(w)

	fyne.Window.CenterOnScreen(w)
	fyne.SetCurrentApp(a)
	w.ShowAndRun()
}

func mmfi(tempbool bool) {
	if tempbool {
		var file string = "c-sharp\\theemployee.txt"
		var tempstr1 string = "1232|Ali|Akbari|09121234567\n11232354534|moslem|moslemi|0999992399\n"
	tmgotocr2:
		file1, err1 := os.Create(file)
		if err1 != nil {
			goto tmgotocr2
		}
		file1.WriteString(tempstr1)

		file1.Close()
	}
}
func insertdepay() {
	var mfile string = "c-sharp\\theemployee.txt"
	months := []string{"farvardin", "ordibehesht", "khordad", "tir", "mordad", "shahrivar", "mehr", "aban", "azar", "dey", "bahman", "esfand"}
tmgotoerr3:
	f, err := os.Open(mfile)
	if err != nil {
		goto tmgotoerr3
	}
	scanner := bufio.NewScanner(f)
	var mainstr1 []string

	for scanner.Scan() {
		var tempstr2 string = (scanner.Text())
		for temp1 := 0; temp1 < len(tempstr2); temp1++ {
			if tempstr2[temp1] == '|' {
				mainstr1 = append(mainstr1, string(scanner.Text()[:temp1]))

				break
			}
		}
	}

	for i := 0; i < len(months); i++ {
		s, _ := os.OpenFile("c-sharp\\payment\\"+months[i]+".txt", os.O_RDWR|os.O_CREATE, 0644)

		scanner1 := bufio.NewScanner(s)
		var mainstr2 [999]string
		var mainstr22 [999]string
		var tempint1 int = 0
		for scanner1.Scan() {

			var tempstr2 string = (scanner1.Text())
			for temp1 := 0; temp1 < len(tempstr2); temp1++ {
				if tempstr2[temp1] == '|' {
					mainstr2[tempint1] = scanner1.Text()[:temp1]
					mainstr22[tempint1] = scanner1.Text()[temp1+1:]
					tempint1++
				}
			}
		}
		for temp1 := 0; temp1 < len(mainstr1); temp1++ {
			if mainstr2[temp1] == "" {
				mainstr2[temp1] = mainstr1[temp1]
				mainstr22[temp1] = "0"
			} else if mainstr1[temp1] != mainstr2[temp1] {
				mainstr2[temp1] = mainstr1[temp1]
				mainstr22[temp1] = "0"
			}
		}
		tempint1 = len(mainstr1) - 1
		strtemp1 := ""
	tmgotocr3:
		s1, err2 := os.Create("c-sharp\\payment\\" + months[i] + ".txt")
		if err2 != nil {
			goto tmgotocr3
		}
		for temp1 := 0; temp1 <= tempint1; temp1++ {
			if mainstr2[temp1] != "" {
				strtemp1 = strtemp1 + mainstr2[temp1] + "|" + mainstr22[temp1] + "\n"
			}
		}
		s1.WriteString(strtemp1)
		s1.Close()
		defer s.Close()
	}
	f.Close()
}
func insertdelea() {
	var mfile string = "c-sharp\\theemployee.txt"
	months := []string{"farvardin", "ordibehesht", "khordad", "tir", "mordad", "shahrivar", "mehr", "aban", "azar", "dey", "bahman", "esfand"}
tgoto1:
	f, err := os.Open(mfile)
	if err != nil {
		goto tgoto1
	}
	scanner := bufio.NewScanner(f)
	var mainstr1 []string

	for scanner.Scan() {
		var tempstr2 string = (scanner.Text())
		for temp1 := 0; temp1 < len(tempstr2); temp1++ {
			if tempstr2[temp1] == '|' {
				mainstr1 = append(mainstr1, string(scanner.Text()[:temp1]))

				break
			}
		}
	}

	for i := 0; i < len(months); i++ {
		s, _ := os.OpenFile("c-sharp\\leave\\"+months[i]+".txt", os.O_RDWR|os.O_CREATE, 0644)
		defer s.Close()
		scanner1 := bufio.NewScanner(s)
		var mainstr2 [999]string
		var mainstr22 [999]string
		var tempint1 int = 0
		for scanner1.Scan() {

			var tempstr2 string = (scanner1.Text())
			for temp1 := 0; temp1 < len(tempstr2); temp1++ {
				if tempstr2[temp1] == '|' {
					mainstr2[tempint1] = scanner1.Text()[:temp1]
					mainstr22[tempint1] = scanner1.Text()[temp1+1:]
					tempint1++
				}
			}
		}
		for temp1 := 0; temp1 < len(mainstr1); temp1++ {
			if mainstr2[temp1] == "" {
				mainstr2[temp1] = mainstr1[temp1]
				mainstr22[temp1] = "0"
			} else if mainstr1[temp1] != mainstr2[temp1] {
				mainstr2[temp1] = mainstr1[temp1]
				mainstr22[temp1] = "0"
			}
		}
		tempint1 = len(mainstr1) - 1
		strtemp1 := ""
	tmgotocr4:
		s1, err4 := os.Create("c-sharp\\leave\\" + months[i] + ".txt")
		if err4 != nil {
			goto tmgotocr4
		}
		for temp1 := 0; temp1 <= tempint1; temp1++ {
			if mainstr2[temp1] != "" {
				strtemp1 = strtemp1 + mainstr2[temp1] + "|" + mainstr22[temp1] + "\n"
			}
		}
		s1.WriteString(strtemp1)
		s1.Close()

	}
	f.Close()
}
