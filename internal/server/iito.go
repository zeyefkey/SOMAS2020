package server

import (
	"github.com/SOMAS2020/SOMAS2020/internal/common"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

// runIITO : IITO makes recommendations about the optimal (and fairest) contributions this term
// to mitigate the common pool dilemma
func (s *SOMASServer) runIITO() error {
	s.logf("start runIITO")
	s.runGiftSession()
	defer s.logf("finish runIITO")
	// TODO:- IITO team
	return nil
}

func (s *SOMASServer) runIITOEndOfTurn() error {
	s.logf("start runIITOEndOfTurn")
	defer s.logf("finish runIITOEndOfTurn")
	// TODO:- IITO team
	return nil
}

func (s *SOMASServer) runGiftSession() ([]common.Action, error) {
	s.logf("start runGiftSession")
	giftRequestDict, err := s.getGiftRequests()
	if err != nil {
		return nil, err
	}
	giftOffersDict, err := s.getGiftOffers(giftRequestDict)
	if err != nil {
		return nil, err
	}
	giftHistoryDict, err := s.getGiftAcceptance(giftOffersDict)
	if err != nil {
		return nil, err
	}
	err = s.distributeGiftHistory(giftHistoryDict)
	if err != nil {
		return nil, err
	}
	// Process actions
	for key, value := range giftHistoryDict {
		s.logf("Gifts from %s: %v\n", key, value)
	}
	defer s.logf("finish runGiftSession")
	return nil, nil
}

func (s *SOMASServer) getGiftRequests() (shared.GiftDict, error) {
	giftRequestDict := shared.GiftDict{}
	var err error
	for id, client := range s.clientMap {
		giftRequestDict[id], err = client.RequestGift()
		if err != nil {
			return giftRequestDict, err
		}
	}
	return giftRequestDict, nil
}
func (s *SOMASServer) getGiftOffers(giftRequestDict shared.GiftDict) (map[shared.ClientID]shared.GiftDict, error) {
	giftOfferDict := map[shared.ClientID]shared.GiftDict{}
	var err error
	for id, client := range s.clientMap {
		giftOfferDict[id], err = client.OfferGifts(giftRequestDict)
		if err != nil {
			return giftOfferDict, err
		}
	}
	return giftOfferDict, nil
}
func (s *SOMASServer) getGiftAcceptance(giftOffersDict map[shared.ClientID]shared.GiftDict) (map[shared.ClientID]shared.GiftInfoDict, error) {
	acceptedGifts := map[shared.ClientID]shared.GiftInfoDict{}
	var err error
	for id, client := range s.clientMap {
		acceptedGifts[id], err = client.AcceptGifts(giftOffersDict[id])
		if err != nil {
			return acceptedGifts, err
		}
	}
	return acceptedGifts, nil
}

func (s *SOMASServer) distributeGiftHistory(acceptedGifts map[shared.ClientID]shared.GiftInfoDict) error {
	//Process acceptedGifts
	for id, client := range s.clientMap {
		err := client.UpdateGiftInfo(acceptedGifts[id])
		if err != nil {
			return err
		}
	}
	return nil
}
