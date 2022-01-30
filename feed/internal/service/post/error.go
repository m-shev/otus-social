package post

const cantMarshalPostIdsErrorTxt = "can't marshal to json post list for user %s, cause: %s\n"
const cantUnmarshalPostIdsErrorTxt = "can't unmarshal from string post lists for user %s, cause: %s\n"
const cantSetPostIdsErrorTxt = "can't set post ids list to cache for user %s, cause: %s\n"
const cantGetPostIDsErrorTxt = "can't get post ids list from cache for user %s cause: %s\n"

func (s *Service) logCantMarshalPostIds(key string, err error) {
	s.logger.Printf(cantMarshalPostIdsErrorTxt, key, err.Error())
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
