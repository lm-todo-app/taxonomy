package resources

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lm-todo-app/taxonomy/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func getTaxonomies(c *fiber.Ctx) error {
	var t models.TaxonomySlice
	var err error

	q := QueryModSlice{}
	q = handleGetTaxonomyParams(c, q)
	q = handlePagination(c, q)

	t, err = models.Taxonomies(q).AllG(c.UserContext())

	if t == nil {
		return fiber.ErrNotFound
	}
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(t)
}

func getTaxonomy(c *fiber.Ctx) error {
	t, err := getTaxonomyFromID(c)
	if err != nil {
		return err
	}
	return c.JSON(t)
}

func createTaxonomy(c *fiber.Ctx) error {
	t, err := getRequestBodyTaxonomy(c)
	if err != nil {
		return err
	}
	err = t.InsertG(c.UserContext(), boil.Infer())
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(t)
}

func updateTaxonomy(c *fiber.Ctx) error {
	new_t, err := getRequestBodyTaxonomy(c)
	if err != nil {
		return err
	}
	t, err := getTaxonomyFromID(c)
	if err != nil {
		return err
	}
	t.Key = new_t.Key
	t.LanguageKey = new_t.LanguageKey
	t.Value = new_t.Value

	_, err = t.UpdateG(c.UserContext(), boil.Infer())
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(t)
}

func deleteTaxonomy(c *fiber.Ctx) error {
	t, err := getTaxonomyFromID(c)
	if err != nil {
		return err
	}
	_, err = t.DeleteG(c.UserContext())
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(t)
}

// func createTaxonomies(c *fiber.Ctx) error {
// }

func getTaxonomyFromID(c *fiber.Ctx) (*models.Taxonomy, error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	t, err := models.FindTaxonomyG(c.UserContext(), id)

	if t == nil && err != nil {
		return nil, fiber.ErrNotFound
	}

	if err != nil {
		return nil, fiber.ErrInternalServerError
	}
	return t, nil
}

func getRequestBodyTaxonomy(c *fiber.Ctx) (*models.Taxonomy, error) {
	new_t := new(models.Taxonomy)
	if err := c.BodyParser(new_t); err != nil {
		return nil, fiber.ErrBadRequest
	}
	return new_t, nil
}
