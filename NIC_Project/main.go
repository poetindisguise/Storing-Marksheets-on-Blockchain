package main

import (
	"NIC-Project/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)
func find(c *gin.Context) {
	roll := c.Query("rollno")
	output, err:= exec.Command("./find.sh", roll).Output()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"err": err})
	}
	c.JSON(http.StatusOK,  string(output))


}
func add(c *gin.Context) {

	var b= make([]byte,c.Request.ContentLength)
	c.Request.Body.Read(b)
    c.Request.Body.Close()
	var rec map[string]string
	json.Unmarshal(b,&rec)
	name:=rec["name"]
	roll:=rec["rollno"]
	year:=rec["year"]
	board:=rec["board"]
	mark:=rec["mark"]
	output, err := exec.Command("./add.sh",name, year, board, mark, roll).Output()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
	}
	c.JSON(http.StatusOK,  string(output))
}
func modify(c *gin.Context) {
	var b= make([]byte,c.Request.ContentLength)
	c.Request.Body.Read(b)

	c.Request.Body.Close()
	var rec map[string]string
	json.Unmarshal(b,&rec)
	roll:=rec["rollno"]
	mark:=rec["mark"]
	output, err := exec.Command("./update.sh",roll,mark).Output()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
	}
	c.JSON(http.StatusOK,  string(output))
}
func all(c *gin.Context) {
	output, err := exec.Command("./AllRecords.sh").Output()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
	}
	c.JSON(http.StatusOK,  string(output))
}
func view(c *gin.Context)  {
	c.Request.ParseForm()
	var students []models.Student
	b := c.Request.FormValue("BOARD")
	e := c.Request.FormValue("EXAM")
	s := c.Request.FormValue("SCHOOL")
	y := c.Request.FormValue("YEAR")
	var names []string
	var rollno []string
	models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_name", &names)
	models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_roll_no", &rollno)
	type TableData struct {
		Name   string
		RollNo string
	}
	data := []TableData{}
	for i := 0; i < len(rollno); i++ {
		data = append(data, TableData{Name: names[i], RollNo: rollno[i]})
	}
	c.HTML(http.StatusOK, "view.html", gin.H{
		"board":  b,
		"school": s,
		"exam":   e,
		"year":   y,
		"name":   names,
		"rollno": rollno,
		"data":   data,
	})

}
func student(c *gin.Context)  {
	var students []models.Student
	var boards []string
	var exam []string
	models.DB.Find(&students).Group("board_name").Pluck("board_name", &boards)
	//var boards = models.DB.Raw("SELECT DISTINCT board_name FROM students")
	models.DB.Find(&students).Group("examination_name").Pluck("examination_name", &exam)
	c.HTML(http.StatusOK, "student.html", gin.H{
		"board": boards,
		"exam":  exam,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("website/*.html")
	//	r.Static("../website/css", "../website/css")
	//	r.Static("website/css/admin.css", "../NIC_Project/website/css/admin.css")

	models.ConnectDataBase()
	r.GET("/student.html", student)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})
	r.GET("/admin", func(c *gin.Context) {
		var students []models.Student
		var boards []string
		var exam []string
		var school []string
		var year []uint

		models.DB.Find(&students).Group("board_name").Pluck("board_name", &boards)
		//var boards = models.DB.Raw("SELECT DISTINCT board_name FROM students")
		models.DB.Find(&students).Group("examination_name").Pluck("examination_name", &exam)
		models.DB.Find(&students).Group("school_name").Pluck("school_name", &school)
		models.DB.Find(&students).Group("year_of_exam").Pluck("year_of_exam", &year)
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"board":  boards,
			"exam":   exam,
			"school": school,
			"year":   year,
		})
	})
	//http.HandleFunc("/displayTable", displayTable)

	r.POST("/add", func(c *gin.Context) {
		c.Request.ParseForm()
		var students []models.Student
		b := c.Request.FormValue("BOARD")
		e := c.Request.FormValue("EXAM")
		s := c.Request.FormValue("SCHOOL")
		y := c.Request.FormValue("YEAR")
		var names []string
		var rollno []string
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_name", &names)
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_roll_no", &rollno)
		type TableData struct {
			Name   string
			RollNo string
		}
		data := []TableData{}
		for i := 0; i < len(rollno); i++ {
			data = append(data, TableData{Name: names[i], RollNo: rollno[i]})
		}
		c.HTML(http.StatusOK, "add.html", gin.H{
			"board":  b,
			"school": s,
			"exam":   e,
			"year":   y,
			"name":   names,
			"rollno": rollno,
			"data":   data,
		})
		//fmt.Fprint(c.Writer,b,e,s,y)

	})
	r.POST("/modify", func(c *gin.Context) {
		c.Request.ParseForm()
		var students []models.Student
		b := c.Request.FormValue("BOARD")
		e := c.Request.FormValue("EXAM")
		s := c.Request.FormValue("SCHOOL")
		y := c.Request.FormValue("YEAR")
		var names []string
		var rollno []string
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_name", &names)
		models.DB.Where("board_name=? AND examination_name=? AND school_name=? AND year_of_exam=?", b, e, s, y).Find(&students).Pluck("student_roll_no", &rollno)
		type TableData struct {
			Name   string
			RollNo string
		}
		data := []TableData{}
		for i := 0; i < len(rollno); i++ {
			data = append(data, TableData{Name: names[i], RollNo: rollno[i]})
		}
		c.HTML(http.StatusOK, "modify.html", gin.H{
			"board":  b,
			"school": s,
			"exam":   e,
			"year":   y,
			"name":   names,
			"rollno": rollno,
			"data":   data,
		})
		//fmt.Fprint(c.Writer,b,e,s,y)

	})

	r.POST("/view",view)
	r.POST("/add_record",add)
	r.POST("modify_record",modify)
	r.GET("/find/",find )
	r.Run()
}
