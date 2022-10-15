package resources

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lm-todo-app/taxonomy/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type QueryModSlice []QueryMod

func (s QueryModSlice) Apply(q *queries.Query) {
	Apply(q, s...)
}

func handleGetTaxonomyParams(c *fiber.Ctx, q QueryModSlice) QueryModSlice {
	if language := c.Query("language"); language != "" {
		fmt.Println(language)
		q = append(q, models.TaxonomyWhere.LanguageKey.EQ(language))
	}
	return q
}

func handlePagination(c *fiber.Ctx, q QueryModSlice) QueryModSlice {
	// TODO Handle errors
	if c.Query("size") == "" && c.Query("page") == "" {
		return q
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		log.Println(err)
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Println(err)
	}
	if size != 0 && page != 0 {
		q = append(q, Limit(size))
		q = append(q, Offset((page-1)*size))
	}
	return q
}

// TODO: Handle order
