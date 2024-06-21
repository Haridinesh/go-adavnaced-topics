package service

import (
	"blogpost/models"
	"blogpost/repository"
)

func Toblogpost(posts []models.Posts) ([]models.PostsResponse, error) {
	postresp := make([]models.PostsResponse, len(posts))
	for i, v := range posts {
		cmnts, err := repository.Dbn.ToGetPostComment(int(v.ID))

		if err != nil {
			return nil, err
		}
		postresp[i].Title = v.Title
		postresp[i].Content = v.Content
		postresp[i].Excerpt = v.Excerpt
		postresp[i].CommentContent = cmnts
		for _, value := range v.Categoriesid {
			category, err := repository.Dbn.ToGetCategoryByID(int(value))

			if err != nil {
				return nil, err
			}
			postresp[i].CategoryName = append(postresp[i].CategoryName, category)
		}
	}
	return postresp, nil
}
