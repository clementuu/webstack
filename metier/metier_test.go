package metier

import (
	"strings"
	"testing"
	"webstack/config"
	"webstack/data"
	"webstack/models"
)

func TestInit(t *testing.T) {
	step1, _ := data.OpenDb(config.GetConfig())
	got := Init(step1)

	if got != nil {
		t.Errorf("got %q, wanted nil", got)
	}
}

type fakeDb struct {
}

func (f fakeDb) AddTodoDb(td models.Todo) error {
	return nil
}
func (f fakeDb) DeleteTodoDb(td models.Todo) error {
	return nil
}
func (f fakeDb) ModifyTodoDb(td models.Todo) error {
	return nil
}
func (f fakeDb) GetTodosDb() (t []models.Todo, e error) {
	return t, nil
}

func TestGetTodos(t *testing.T) {
	db := fakeDb{}
	Init(db)

	want := []models.Todo{}
	got, err := GetTodos()

	if err != nil {
		t.Errorf("Erreur lors de la récupération des données : %v", err)
	}
	if len(got) != len(want) {
		t.Errorf("Expected %d items, but got %d items", len(want), len(got))
		return
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Expected item %d to be '%v', but got '%v'", i, want[i], got[i])
		}
	}
}

func TestAddTodo(t *testing.T) {
	db := fakeDb{}
	Init(db)

	var tests = []struct {
		name, entryTxt, want string
	}{
		{"Cas normal", "Blablabla", "Blablabla"},
		{"Chaîne vide", "", "veuillez renseigner du texte"},
		{"Caractères spéciaux autorisés", "(/$-_~+)=", "(/$-_~+)="},
		{"Caractères spéciaux non autorisés", "(/$-_]&[~]%)=", "caractères spéciaux non autorisés"},
		{"Plusieurs espaces en entrée", "    ", "veuillez renseigner du texte"},
		{"Chaîne longue", "Une chaine très longue mais sans caractères spéciaux, d'ailleurs ma mère me dit toujours que je suis spécial, ça va c'est assez long ? Bon aller on va dire que oui", "Une chaine très longue mais sans caractères spéciaux, d'ailleurs ma mère me dit toujours que je suis spécial, ça va c'est assez long ? Bon aller on va dire que oui"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddTodo(tt.entryTxt)
			if !strings.Contains(got, tt.want) && !strings.Contains(err.Error(), tt.want) {
				t.Errorf("expected response to contain '%s', but got '%s'", tt.want, err.Error())
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	db := fakeDb{}
	Init(db)

	var tests = []struct {
		name, entryTxt, entryId, want string
	}{
		{"Cas normal", "Blablabla", "3", "Blablabla"},
		{"Chaîne vide", "", "123", "réessayez ultérieurement"},
		{"Id non numérique", "BlablaASupprimer", "azerty", "erreur de conversion"},
		{"Id vide", "BlablaASupprimer2", "", "réessayez ultérieurement"},
		{"Chaîne longue", "Une chaine très longue mais sans caractères spéciaux, d'ailleurs ma mère me dit toujours que je suis spécial, ça va c'est assez long ? Bon aller on va dire que oui", "10", "Une chaine très longue mais sans caractères spéciaux, d'ailleurs ma mère me dit toujours que je suis spécial, ça va c'est assez long ? Bon aller on va dire que oui"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteTodo(tt.entryId, tt.entryTxt)
			if !strings.Contains(got, tt.want) && !strings.Contains(err.Error(), tt.want) {
				t.Errorf("expected response to contain '%s', but got '%s'", tt.want, err.Error())
			}
		})
	}
}

func TestModifyTodo(t *testing.T) {
	db := fakeDb{}
	Init(db)

	var tests = []struct {
		name, entryTxt, entryId, want string
	}{
		{"Cas normal", "Blabliblou", "3", "Blabliblou"},
		{"Chaîne vide", "", "123", "veuillez renseigner du texte"},
		{"Caractères spéciaux autorisés", "(/$-_~+)=", "13", "(/$-_~+)="},
		{"Caractères spéciaux non autorisés", "(/${}-_~+)=", "13", "caractères spéciaux non autorisés"},
		{"Id non numérique", "BlablaAModifier", "azerty", "erreur de conversion"},
		{"Id vide", "BlablaAModifier2", "", "réessayez ultérieurement"},
		{"Plusieurs espaces en entrée", "    ", "56", "veuillez renseigner du texte"},
		{"Chaîne longue", "Une chaine très longue mais sans caractères spéciaux, d'ailleurs ma mère me dit toujours que je suis spécial, ça va c'est assez long ? Bon aller on va dire que oui", "2", "Une chaine très longue mais sans caractères spéciaux, d'ailleurs ma mère me dit toujours que je suis spécial, ça va c'est assez long ? Bon aller on va dire que oui"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ModifyTodo(tt.entryId, tt.entryTxt)
			if !strings.Contains(got, tt.want) && !strings.Contains(err.Error(), tt.want) {
				t.Errorf("expected response to contain '%s', but got '%s'", tt.want, err.Error())
			}
		})
	}
}