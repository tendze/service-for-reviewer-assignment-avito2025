package response

import "dang.z.v.task/internal/domain"

type AddTeamInfoResponse struct {
	TeamName string       `json:"team_name"`
	Members  []TeamMember `json:"members"`
}

type TeamMember struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func NewAddTeamResponse(teamName string, teamMembers *[]domain.User) map[string]interface{} {
	return map[string]interface{}{
		"team": usersToTeamMembers(teamMembers),
	}
}

type GetTeamResponse struct {
	TeamName string       `json:"team_name"`
	Members  []TeamMember `json:"members"`
}

func NewGetTeamResponse(teamName string, teamMembers *[]domain.User) GetTeamResponse {
	return GetTeamResponse{
		TeamName: teamName,
		Members:  usersToTeamMembers(teamMembers),
	}
}

func usersToTeamMembers(teamMembers *[]domain.User) []TeamMember {
	if teamMembers == nil {
		return []TeamMember{}
	}

	res := make([]TeamMember, 0, len(*teamMembers))
	for _, user := range *teamMembers {
		res = append(
			res,
			TeamMember{
				UserID:   user.ID,
				Username: user.Name,
				IsActive: user.IsActive,
			},
		)
	}

	return res
}
