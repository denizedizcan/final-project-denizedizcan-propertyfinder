package handlers

/*
func (h handler) handleValidateToken(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	token := r.Header.Get("Authorization")

	if token == "" {
		return errors.New("Empty Token")
	}

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		return err
	}

	user.Prepare()

	if err := user.ValidateToken(h.DB, token); err != nil {
		return err
	}
	return nil
}
*/
