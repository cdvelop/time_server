package timeserver

func (t *timeServer) SetDate(new_date string) {
	t.fake_date = new_date
}
