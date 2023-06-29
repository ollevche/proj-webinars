package main

type Storage struct {
	wallets []*Wallet
}

func NewStorage() *Storage {
	return &Storage{
		wallets: []*Wallet{
			{
				ID:         "649dbb5cd7c795f157361bc4",
				HolderName: "John",
				Amount:     10,
			},
			{
				ID:         "649dbb648ab17759e1dd6546",
				HolderName: "Peter",
				Amount:     0,
			},
		},
	}
}

func (s *Storage) GetAllWallets() []*Wallet {
	return s.wallets
}

func (s *Storage) GetWalletByID(id string) (*Wallet, bool) {
	for _, w := range s.wallets {
		if w.ID == id {
			return w, true
		}
	}

	return nil, false
}

func (s *Storage) InsertWallet(w *Wallet) {
	s.wallets = append(s.wallets, w)
}
