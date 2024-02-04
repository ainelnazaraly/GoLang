package characters 
import (
	
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type Character struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	IsRat      bool   `json:"isRat"`
}

var CharactersALl = []Character{
    {1, "Katniss", "The Mockingjay", true},
    {2, "Peeta", "The Baker", true},
    {3, "Gale", "The Rebel", true},
    {4, "Haymitch", "The Mentor", true},
    {5, "Effie", "The Escort", true},
    {6, "President Snow", "The President of Panem", true},
    {7, "Finnick", "The Victor", true},
    {8, "Johanna", "The Survivor", true},
    {9, "Cinna", "The Stylist", true},
    {10, "Primrose", "The Healer", true},
    {11, "Rue", "The Tribute", false},
}
func GetCharacterList() []Character {
    return CharactersALl
}

func GetCharacterByNameInternal(name string) *Character {
	for _, c := range CharactersALl {
		if c.Name == name {
			return &c
		}
	}
	return nil
}
type Response struct {
	Characters []Character `json:"persons"`
}
func prepareCharacterList() []Character {
	var characters []Character

	characters = append(characters, Character{ID: 1, Name: "Katniss", Role: "The Mockingjay", IsRat: false})
	characters = append(characters, Character{ID: 2, Name: "Peeta", Role: "The Baker", IsRat: false})
	characters = append(characters, Character{ID: 3, Name: "Gale", Role: "The Rebel", IsRat: false})
	characters = append(characters, Character{ID: 4, Name: "Haymitch", Role: "The Mentor", IsRat: false})
	characters = append(characters, Character{ID: 5, Name: "Effie", Role: "The Escort", IsRat: false})
	characters = append(characters, Character{ID: 6, Name: "President Snow", Role: "The President of Panem", IsRat: false})
	characters = append(characters, Character{ID: 7, Name: "Finnick", Role: "The Victor", IsRat: false})
	characters = append(characters, Character{ID: 8, Name: "Johanna", Role: "The Survivor", IsRat: false})
	characters = append(characters, Character{ID: 9, Name: "Cinna", Role: "The Stylist", IsRat: false})
	characters = append(characters, Character{ID: 10, Name: "Primrose", Role: "The Healer", IsRat: false})

	return characters
}

func Characters(w http.ResponseWriter, r *http.Request) {
	var response Response
	characters := GetCharacterList()
	response.Characters = characters
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func GetCharacters(w http.ResponseWriter, r *http.Request) {

	charactersList := GetCharacterList()

	jsonResponse, err := json.Marshal(charactersList)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetCharacterByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	character := GetCharacterByNameInternal(name)

	if character == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(character)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}



