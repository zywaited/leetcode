package one

type authenticationNode struct {
	tokenId string
	expire  int
}

type AuthenticationManager struct {
	tokenIdExpires map[string]int
	tokenIds       []authenticationNode
	timeToLive     int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{tokenIdExpires: map[string]int{}, timeToLive: timeToLive}
}

func (authentication *AuthenticationManager) Generate(tokenId string, currentTime int) {
	expire := currentTime + authentication.timeToLive
	authentication.tokenIdExpires[tokenId] = expire
	authentication.tokenIds = append(authentication.tokenIds, authenticationNode{tokenId: tokenId, expire: expire})
}

func (authentication *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if currentTime < authentication.tokenIdExpires[tokenId] {
		authentication.Generate(tokenId, currentTime)
	}
}

func (authentication *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	index := 0
	for ; index < len(authentication.tokenIds) && authentication.tokenIds[index].expire <= currentTime; index++ {
		if authentication.tokenIds[index].expire == authentication.tokenIdExpires[authentication.tokenIds[index].tokenId] {
			delete(authentication.tokenIdExpires, authentication.tokenIds[index].tokenId)
		}
	}
	authentication.tokenIds = authentication.tokenIds[index:]
	return len(authentication.tokenIdExpires)
}
