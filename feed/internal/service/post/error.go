package post

const cantMarshalPostIdsErrorTxt = "can't marshal to json post list for user %s, cause: %s\n"
const cantMarshalPostErrorTxt = "can't marshal post with id -> %d cause: %s\n"
const cantUnmarshalPostErrorTxt = "can't unmarshal post with id -> %d cause: %s\n"
const cantUnmarshalPostIdsErrorTxt = "can't unmarshal from string post lists for user %s, cause: %s\n"
const cantSetPostIdsErrorTxt = "can't set post ids list to consumer cache for user %s, cause: %s\n"
const cantSetPostErrorTxt = "can't set post with id -> %d to post cache cause: %s\n"
const cantGetPostIDsErrorTxt = "can't get post ids list from consumer cache for user %s cause: %s\n"
const cantGetPostErrorTxt = "can't get post with id -> %d from post cache cause: %s\n"
const cantDelPostErrorText = "can't delete post with id -> %d from post cache cause: %'"

func (s *Service) logCantMarshalPostIds(key string, err error) {
	s.logger.Printf(cantMarshalPostIdsErrorTxt, key, err.Error())
}

func (s *Service) logCantMarshalPost(id int, err error) {
	s.logger.Printf(cantMarshalPostErrorTxt, id, err.Error())
}

func (s *Service) logCantUnmarshalPost(id int, err error) {
	s.logger.Printf(cantUnmarshalPostErrorTxt, id, err.Error())
}

func (s *Service) logCantUnmarshalPostIds(key string, err error) {
	s.logger.Printf(cantUnmarshalPostIdsErrorTxt, key, err.Error())
}

func (s *Service) logCantSetPostIds(key string, err error) {
	s.logger.Printf(cantSetPostIdsErrorTxt, key, err.Error())
}

func (s *Service) logCantGetPostIds(key string, err error) {
	s.logger.Printf(cantGetPostIDsErrorTxt, key, err.Error())
}

func (s *Service) logCantGetPost(id int, err error) {
	s.logger.Printf(cantGetPostErrorTxt, id, err.Error())
}

func (s *Service) logCantSetPost(id int, err error) {
	s.logger.Printf(cantSetPostErrorTxt, id, err.Error())
}

func (s *Service) logCantDelPost(id int, err error) {
	s.logger.Printf(cantDelPostErrorText, id, err.Error())
}
