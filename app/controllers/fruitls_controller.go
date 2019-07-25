package controllers

import (
	"github.com/revel/revel"
	"go-revel-rest/app/models"
	"strings"
	//"strconv"
)

type FruitsController struct {
	*revel.Controller
}

func (c FruitsController) GetFruitById(id string) revel.Result {

	// get data
	fruits, err := Dbm.Select(models.Fruit{},
		`SELECT * FROM fruits  where id=`+id)
	if err != nil {
		panic(err)
	}

	response := JsonResponse{}
	 if len(fruits) >0{
		 response.Data = fruits[0]
	 }
	return c.RenderJSON(response)
}

func (c FruitsController) GetFruits() revel.Result {

	// get data
	fruits, err := Dbm.Select(models.Fruit{},
		`SELECT * FROM fruits `)
	if err != nil {
		panic(err)
	}
	response := JsonResponse{}
	response.Data = fruits

	return c.RenderJSON(response)
}
func (c FruitsController) DeleteFruit(id string) revel.Result {

	// delete data
	result, err := Dbm.Exec(`delete FROM fruits  where id=`+id)
	if err != nil {
		panic(err)
	}

	response := JsonResponse{}
	response.Data=result
	return c.RenderJSON(response)
}

func (c FruitsController) SaveFruit(name string,value string) revel.Result {
	var query strings.Builder
	query.WriteString("insert into fruits SET name=")
	query.WriteString("'")
	query.WriteString(name)
	query.WriteString("'")
	query.WriteString(",value=")
	query.WriteString(value)
	//panic(query.String())
	fruits, err := Dbm.Select(models.Fruit{},
		query.String())
	if err != nil {
		//panic(err)
		panic(query.String())
	}

	response := JsonResponse{}
	response.Success = true
	response.Data=fruits
	if err != nil {
		response.Error = err.Error()
	}
	return c.RenderJSON(response)
}

func (c FruitsController) UpdateFruit(id string,name string,value string) revel.Result {

	// get data
	fruit, err := Dbm.Select(models.Fruit{},
		`SELECT * FROM fruits  where id=`+id)
	if err != nil {
		panic(err)
	}

	response := JsonResponse{}
	if len(fruit) >0{
		var query strings.Builder
		query.WriteString("update fruits SET name=")
		query.WriteString("'")
		query.WriteString(name)
		query.WriteString("'")
		query.WriteString(",value=")
		query.WriteString(value)
		query.WriteString(" where id=")
		query.WriteString(id)
		//panic(query.String())
	     Dbm.Exec(query.String())
		if err != nil {
			//panic(err)
			panic(query.String())
		}


		response.Success = true
		if err != nil {
			response.Error = err.Error()
		}

	}	else{
		response.Success = true
		response.Error = "Item not found"

	}


	return c.RenderJSON(response)
}
