package handlers

import (
	"esapp/internal/dto"
	"esapp/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	groupService services.GroupService
}

func NewGroupHandler(groupService services.GroupService) *GroupHandler {
	return &GroupHandler{groupService: groupService}
}

func (h *GroupHandler) CreateGroup(c *gin.Context) {
	val, _ := c.Get("userID")
	userID := val.(uint)

	var req dto.CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.groupService.CreateGroup(userID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "group created"})
}

func (h *GroupHandler) AddMember(c *gin.Context) {
	val, _ := c.Get("userID")
	requestID := val.(uint)

	groupIDParam := c.Param("groupId")
	groupID, _ := strconv.ParseUint(groupIDParam, 10, 64)

	var req dto.AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.groupService.AddMemberByEmail(requestID, uint(groupID), req.Email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "member added"})

}

func (h *GroupHandler) ListGroups(c *gin.Context) {
	val, _ := c.Get("userID")
	userID := val.(uint)

	groups, err := h.groupService.ListGroups(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch groups"})
		return
	}

	var response []dto.GroupResponse

	for _, g := range groups {
		var membersDto []dto.Member
		for _, u := range g.Members {
			membersDto = append(membersDto, dto.Member{
				ID:       u.ID,
				Username: u.Username,
				Email:    u.Email,
			})
		}

		response = append(response, dto.GroupResponse{
			ID:          g.ID,
			Name:        g.Name,
			CreatorName: g.Creator.Username,
			CreatedAt:   g.CreatedAt,
			Members:     membersDto,
		})
	}

	c.JSON(200, response)
}

func (h *GroupHandler) ListMembers(c *gin.Context) {
	val, _ := c.Get("userID")
	requesterID := val.(uint)

	groupIDParam := c.Param("groupId")
	groupID, _ := strconv.ParseUint(groupIDParam, 10, 64)

	members, err := h.groupService.ListMembers(requesterID, uint(groupID))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := make([]gin.H, 0)
	for _, u := range members {
		response = append(response, gin.H{
			"id":       u.ID,
			"username": u.Username,
			"email":    u.Email,
		})
	}

	c.JSON(200, response)
}

func (h *GroupHandler) GetGroupByID(c *gin.Context) {
	val, _ := c.Get("userID")
	userID := val.(uint)

	groupIDParam := c.Param("groupId")
	groupID, _ := strconv.ParseUint(groupIDParam, 10, 64)

	group, err := h.groupService.GetGroupByID(uint(groupID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}
