package main

type postCategory struct {
	IdPostCategory int8   `json:"id_post_category"`
	Category       string `json:"category"`
	CreatedAt      string `json:"created_at"`
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8000")
}

var postCategories = []postCategory{
	{1, "sports", "2022-02-01"},
	{2, "game", "2022-02-02"},
	{3, "music", "2022-02-03"},
}

func getPostCategories(context *gin.Context) {
	context.intendedJSON(http.StatusOk, postCategories)
}
