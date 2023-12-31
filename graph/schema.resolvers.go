package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"go_test/graph/model"
	"go_test/models"
	"strconv"
)

// CreatePerson is the resolver for the createPerson field.
func (r *mutationResolver) CreatePerson(ctx context.Context, person *model.NewPerson) (*model.PersonMutationPayload, error) {
	if person.Age == 0 {
		person.Age = r.App.GeneratorService.GetAgeGeneratorResult(person.Name)
	}
	if len(person.Gender) == 0 {
		person.Gender = r.App.GeneratorService.GetGenderGeneratorResult(person.Name)
	}
	if len(person.Nationality) == 0 {
		person.Nationality = r.App.GeneratorService.GetNationalityGeneratorResult(person.Name)
	}

	dbTypePerson := models.Person{
		Name:        person.Name,
		Surname:     person.Surname,
		Patronymic:  person.Patronymic,
		Age:         person.Age,
		Gender:      person.Gender,
		Nationality: person.Nationality,
	}

	_, err := r.App.DB.SavePerson(dbTypePerson)
	if err != nil {
		r.App.Logger.Info("Cannot save person to database")
		return &model.PersonMutationPayload{Ok: false}, err
	}
	return &model.PersonMutationPayload{Ok: true}, nil
}

// DeletePerson is the resolver for the deletePerson field.
func (r *mutationResolver) DeletePerson(ctx context.Context, id int) (*model.PersonMutationPayload, error) {
	if id != 0 {
		err := r.App.DB.DeletePersonByID(int64(id))
		if err != nil {
			r.App.Logger.Info("Cannot delete person")
			return &model.PersonMutationPayload{Ok: false}, err
		}
	}
	return &model.PersonMutationPayload{Ok: true}, nil
}

// UpdatePerson is the resolver for the updatePerson field.
func (r *mutationResolver) UpdatePerson(ctx context.Context, person model.PersonInput) (*model.PersonMutationPayload, error) {

	id, err := strconv.Atoi(person.ID)
	if err != nil {
		r.App.Logger.Info("Failed to parse personId")
		return &model.PersonMutationPayload{Ok: false}, err
	}
	dbTypePerson := models.Person{
		Name:        person.Name,
		Surname:     person.Surname,
		Patronymic:  person.Patronymic,
		Age:         person.Age,
		Gender:      person.Gender,
		Nationality: person.Nationality,
		Id:          id,
	}

	err = r.App.DB.UpdatePerson(dbTypePerson)

	if err != nil {
		r.App.Logger.Info("cannot update person", err)
		return &model.PersonMutationPayload{Ok: false}, err
	}
	return &model.PersonMutationPayload{Ok: true}, nil
}

// Persons is the resolver for the Persons field.
func (r *queryResolver) Persons(ctx context.Context, limit int, offset int, filter string) ([]*model.Person, error) {
	params := make(map[string]string)
	if int(limit) != 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	if offset != 0 {
		params["offset"] = strconv.Itoa(offset)
	}

	if len(filter) != 0 {
		params["filter"] = filter
	}

	persons, err := r.App.DB.GetPersons(params)
	if err != nil {
		r.App.Logger.Info("error while tried get persons from db")
		return nil, err
	}
	var resPerson []*model.Person
	for _, v := range persons {
		bufPerson := &model.Person{
			ID:          strconv.Itoa(v.Id),
			Name:        v.Name,
			Surname:     v.Surname,
			Patronymic:  v.Patronymic,
			Age:         v.Age,
			Nationality: v.Nationality,
			Gender:      v.Gender,
		}
		resPerson = append(resPerson, bufPerson)
	}
	return resPerson, nil
}

// PersonByID is the resolver for the PersonById field.
func (r *queryResolver) PersonByID(ctx context.Context, id int) (*model.Person, error) {
	person, err := r.App.DB.GetPersonByID(int64(id))
	if err != nil {
		r.App.Logger.Info("Error when tried get person from db")
		return nil, err
	}
	resPerson := &model.Person{
		ID:          strconv.Itoa(person.Id),
		Name:        person.Name,
		Surname:     person.Surname,
		Patronymic:  person.Patronymic,
		Age:         person.Age,
		Nationality: person.Nationality,
		Gender:      person.Gender,
	}
	return resPerson, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
