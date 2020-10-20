package helpers

import (
	"sample-project/app/models/entities"
)

var cach *cache

// ICache интерфейс кэша
type ICache interface {
	TokenIsActualByToken(string) (string, bool)                 // актуальность токена по токену
	TokenIsActualBySID(string) (string, bool)                   // актуальность токена по Session ID
	Get(sid string) (u *entities.User, token string, err error) // Get получение данных кэша по SID
	Set(sid, token string, user *entities.User) (err error)     // Set установка данных кэша по SID
	Delete(sid string) (err error)                              // Delete удаление данных кэша по SID
}

// cache тип кэша
type cache struct {
	actualToken map[string]string         // [:Session ID:]:token:
	actualUsers map[string]*entities.User // [:Session ID:]:user:
}

// GetCache получение экземпляра кэша
func GetCache() (c ICache, err error) {
	if cach == nil {
		cach = new(cache)
		cach.actualToken = make(map[string]string)
		cach.actualUsers = make(map[string]*entities.User)
	}

	return cach, nil
}

// TokenIsActualByToken актуальность токена по токену
func (c *cache) TokenIsActualByToken(token string) (string, bool) {
	for _, t := range c.actualToken {
		if token == t {
			return token, true
		}
	}

	return "", false
}

// TokenIsActualBySID актуальность токена по Session ID
func (c *cache) TokenIsActualBySID(sid string) (token string, ok bool) {
	if token, ok = c.actualToken[sid]; ok {
		return
	}

	return
}

// Get получение данных кэша по SID
func (c *cache) Get(sid string) (u *entities.User, token string, err error) {
	return c.actualUsers[sid], c.actualToken[sid], nil
}

// Set установка данных кэша по SID
func (c *cache) Set(sid, token string, user *entities.User) (err error) {
	c.actualToken[sid] = token
	c.actualUsers[sid] = user

	return
}

// Delete удаление данных кэша по SID
func (c *cache) Delete(sid string) (err error) {
	delete(c.actualToken, sid)
	delete(c.actualUsers, sid)

	return
}
