package chain

import "rate/internal/pkg/provider"

type Chain interface {
	provider.IRateProvider
	SetNext(ProviderNode)
}

type ProviderNode struct {
	provider provider.IRateProvider
	next     *ProviderNode
}

func NewProviderNode(provider provider.IRateProvider) *ProviderNode {
	return &ProviderNode{
		provider: provider,
	}
}

func (pn *ProviderNode) GetRate() (*float64, error) {
	rate, err := pn.provider.GetRate()

	if err != nil {
		if pn.next != nil {
			return pn.next.GetRate()
		} else {
			return nil, err
		}
	}

	return rate, nil
}

func (pn *ProviderNode) SetNext(node *ProviderNode) {
	pn.next = node
}
