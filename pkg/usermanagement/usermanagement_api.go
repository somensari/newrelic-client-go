// Code generated by tutone: DO NOT EDIT
package usermanagement

import "context"

// A mutation for adding user(s) to group(s).
func (a *Usermanagement) UserManagementAddUsersToGroups(
	addUsersToGroupsOptions UserManagementUsersGroupsInput,
) (*UserManagementAddUsersToGroupsPayload, error) {
	return a.UserManagementAddUsersToGroupsWithContext(context.Background(),
		addUsersToGroupsOptions,
	)
}

// A mutation for adding user(s) to group(s).
func (a *Usermanagement) UserManagementAddUsersToGroupsWithContext(
	ctx context.Context,
	addUsersToGroupsOptions UserManagementUsersGroupsInput,
) (*UserManagementAddUsersToGroupsPayload, error) {

	resp := UserManagementAddUsersToGroupsQueryResponse{}
	vars := map[string]interface{}{
		"addUsersToGroupsOptions": addUsersToGroupsOptions,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, UserManagementAddUsersToGroupsMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.UserManagementAddUsersToGroupsPayload, nil
}

type UserManagementAddUsersToGroupsQueryResponse struct {
	UserManagementAddUsersToGroupsPayload UserManagementAddUsersToGroupsPayload `json:"UserManagementAddUsersToGroups"`
}

const UserManagementAddUsersToGroupsMutation = `mutation(
	$addUsersToGroupsOptions: UserManagementUsersGroupsInput,
) { userManagementAddUsersToGroups(
	addUsersToGroupsOptions: $addUsersToGroupsOptions,
) {
	groups {
		displayName
		id
		users {
			nextCursor
			totalCount
		}
	}
} }`

// A mutation for creating a group in an authentication domain.
func (a *Usermanagement) UserManagementCreateGroup(
	createGroupOptions UserManagementCreateGroup,
) (*UserManagementCreateGroupPayload, error) {
	return a.UserManagementCreateGroupWithContext(context.Background(),
		createGroupOptions,
	)
}

// A mutation for creating a group in an authentication domain.
func (a *Usermanagement) UserManagementCreateGroupWithContext(
	ctx context.Context,
	createGroupOptions UserManagementCreateGroup,
) (*UserManagementCreateGroupPayload, error) {

	resp := UserManagementCreateGroupQueryResponse{}
	vars := map[string]interface{}{
		"createGroupOptions": createGroupOptions,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, UserManagementCreateGroupMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.UserManagementCreateGroupPayload, nil
}

type UserManagementCreateGroupQueryResponse struct {
	UserManagementCreateGroupPayload UserManagementCreateGroupPayload `json:"UserManagementCreateGroup"`
}

const UserManagementCreateGroupMutation = `mutation(
	$createGroupOptions: UserManagementCreateGroup,
) { userManagementCreateGroup(
	createGroupOptions: $createGroupOptions,
) {
	group {
		displayName
		id
		users {
			nextCursor
			totalCount
		}
	}
} }`

// A mutation for creating a user in an authentication domain.
func (a *Usermanagement) UserManagementCreateUser(
	createUserOptions UserManagementCreateUser,
) (*UserManagementCreateUserPayload, error) {
	return a.UserManagementCreateUserWithContext(context.Background(),
		createUserOptions,
	)
}

// A mutation for creating a user in an authentication domain.
func (a *Usermanagement) UserManagementCreateUserWithContext(
	ctx context.Context,
	createUserOptions UserManagementCreateUser,
) (*UserManagementCreateUserPayload, error) {

	resp := UserManagementCreateUserQueryResponse{}
	vars := map[string]interface{}{
		"createUserOptions": createUserOptions,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, UserManagementCreateUserMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.UserManagementCreateUserPayload, nil
}

type UserManagementCreateUserQueryResponse struct {
	UserManagementCreateUserPayload UserManagementCreateUserPayload `json:"UserManagementCreateUser"`
}

const UserManagementCreateUserMutation = `mutation(
	$createUserOptions: UserManagementCreateUser!,
) { userManagementCreateUser(
	createUserOptions: $createUserOptions,
) {
	createdUser {
		authenticationDomainId
		email
		id
		name
		type {
			displayName
			id
		}
	}
} }`

// A mutation for deleting a group.
func (a *Usermanagement) UserManagementDeleteGroup(
	groupOptions UserManagementDeleteGroup,
) (*UserManagementDeleteGroupPayload, error) {
	return a.UserManagementDeleteGroupWithContext(context.Background(),
		groupOptions,
	)
}

// A mutation for deleting a group.
func (a *Usermanagement) UserManagementDeleteGroupWithContext(
	ctx context.Context,
	groupOptions UserManagementDeleteGroup,
) (*UserManagementDeleteGroupPayload, error) {

	resp := UserManagementDeleteGroupQueryResponse{}
	vars := map[string]interface{}{
		"groupOptions": groupOptions,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, UserManagementDeleteGroupMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.UserManagementDeleteGroupPayload, nil
}

type UserManagementDeleteGroupQueryResponse struct {
	UserManagementDeleteGroupPayload UserManagementDeleteGroupPayload `json:"UserManagementDeleteGroup"`
}

const UserManagementDeleteGroupMutation = `mutation(
	$groupOptions: UserManagementDeleteGroup,
) { userManagementDeleteGroup(
	groupOptions: $groupOptions,
) {
	group {
		displayName
		id
		users {
			nextCursor
			totalCount
		}
	}
} }`

// A mutation for deleting a user.
func (a *Usermanagement) UserManagementDeleteUser(
	deleteUserOptions UserManagementDeleteUser,
) (*UserManagementDeleteUserPayload, error) {
	return a.UserManagementDeleteUserWithContext(context.Background(),
		deleteUserOptions,
	)
}

// A mutation for deleting a user.
func (a *Usermanagement) UserManagementDeleteUserWithContext(
	ctx context.Context,
	deleteUserOptions UserManagementDeleteUser,
) (*UserManagementDeleteUserPayload, error) {

	resp := UserManagementDeleteUserQueryResponse{}
	vars := map[string]interface{}{
		"deleteUserOptions": deleteUserOptions,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, UserManagementDeleteUserMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.UserManagementDeleteUserPayload, nil
}

type UserManagementDeleteUserQueryResponse struct {
	UserManagementDeleteUserPayload UserManagementDeleteUserPayload `json:"UserManagementDeleteUser"`
}

const UserManagementDeleteUserMutation = `mutation(
	$deleteUserOptions: UserManagementDeleteUser!,
) { userManagementDeleteUser(
	deleteUserOptions: $deleteUserOptions,
) {
	deletedUser {
		id
	}
} }`

// A mutation for removing user(s) from group(s).
func (a *Usermanagement) UserManagementRemoveUsersFromGroups(
	removeUsersFromGroupsOptions UserManagementUsersGroupsInput,
) (*UserManagementRemoveUsersFromGroupsPayload, error) {
	return a.UserManagementRemoveUsersFromGroupsWithContext(context.Background(),
		removeUsersFromGroupsOptions,
	)
}

// A mutation for removing user(s) from group(s).
func (a *Usermanagement) UserManagementRemoveUsersFromGroupsWithContext(
	ctx context.Context,
	removeUsersFromGroupsOptions UserManagementUsersGroupsInput,
) (*UserManagementRemoveUsersFromGroupsPayload, error) {

	resp := UserManagementRemoveUsersFromGroupsQueryResponse{}
	vars := map[string]interface{}{
		"removeUsersFromGroupsOptions": removeUsersFromGroupsOptions,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, UserManagementRemoveUsersFromGroupsMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.UserManagementRemoveUsersFromGroupsPayload, nil
}

type UserManagementRemoveUsersFromGroupsQueryResponse struct {
	UserManagementRemoveUsersFromGroupsPayload UserManagementRemoveUsersFromGroupsPayload `json:"UserManagementRemoveUsersFromGroups"`
}

const UserManagementRemoveUsersFromGroupsMutation = `mutation(
	$removeUsersFromGroupsOptions: UserManagementUsersGroupsInput!,
) { userManagementRemoveUsersFromGroups(
	removeUsersFromGroupsOptions: $removeUsersFromGroupsOptions,
) {
	groups {
		displayName
		id
		users {
			nextCursor
			totalCount
		}
	}
} }`

// A mutation for updating an existing group.
func (a *Usermanagement) UserManagementUpdateGroup(
	updateGroupOptions UserManagementUpdateGroup,
) (*UserManagementUpdateGroupPayload, error) {
	return a.UserManagementUpdateGroupWithContext(context.Background(),
		updateGroupOptions,
	)
}

// A mutation for updating an existing group.
func (a *Usermanagement) UserManagementUpdateGroupWithContext(
	ctx context.Context,
	updateGroupOptions UserManagementUpdateGroup,
) (*UserManagementUpdateGroupPayload, error) {

	resp := UserManagementUpdateGroupQueryResponse{}
	vars := map[string]interface{}{
		"updateGroupOptions": updateGroupOptions,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, UserManagementUpdateGroupMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.UserManagementUpdateGroupPayload, nil
}

type UserManagementUpdateGroupQueryResponse struct {
	UserManagementUpdateGroupPayload UserManagementUpdateGroupPayload `json:"UserManagementUpdateGroup"`
}

const UserManagementUpdateGroupMutation = `mutation(
	$updateGroupOptions: UserManagementUpdateGroup,
) { userManagementUpdateGroup(
	updateGroupOptions: $updateGroupOptions,
) {
	group {
		displayName
		id
		users {
			nextCursor
			totalCount
		}
	}
} }`

// A mutation for updating an existing user.
func (a *Usermanagement) UserManagementUpdateUser(
	updateUserOptions UserManagementUpdateUser,
) (*UserManagementUpdateUserPayload, error) {
	return a.UserManagementUpdateUserWithContext(context.Background(),
		updateUserOptions,
	)
}

// A mutation for updating an existing user.
func (a *Usermanagement) UserManagementUpdateUserWithContext(
	ctx context.Context,
	updateUserOptions UserManagementUpdateUser,
) (*UserManagementUpdateUserPayload, error) {

	resp := UserManagementUpdateUserQueryResponse{}
	vars := map[string]interface{}{
		"updateUserOptions": updateUserOptions,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, UserManagementUpdateUserMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.UserManagementUpdateUserPayload, nil
}

type UserManagementUpdateUserQueryResponse struct {
	UserManagementUpdateUserPayload UserManagementUpdateUserPayload `json:"UserManagementUpdateUser"`
}

const UserManagementUpdateUserMutation = `mutation(
	$updateUserOptions: UserManagementUpdateUser!,
) { userManagementUpdateUser(
	updateUserOptions: $updateUserOptions,
) {
	user {
		email
		emailVerificationState
		groups {
			nextCursor
			totalCount
		}
		id
		lastActive
		name
		pendingUpgradeRequest {
			id
			message
		}
		timeZone
		type {
			displayName
			id
		}
	}
} }`
