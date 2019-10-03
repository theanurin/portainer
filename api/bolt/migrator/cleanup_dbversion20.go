package migrator

import portainer "github.com/portainer/portainer/api"

func (m *Migrator) cleanupAccessPolicies() error {
	endpoints, err := m.endpointService.Endpoints()
	if err != nil {
		return err
	}

	for _, endpoint := range endpoints {
		for policyID := range endpoint.UserAccessPolicies {
			_, err := m.userService.User(policyID)
			if err == portainer.ErrObjectNotFound {
				delete(endpoint.UserAccessPolicies, policyID)
			}
		}

		for policyID := range endpoint.TeamAccessPolicies {
			_, err := m.teamService.Team(policyID)
			if err == portainer.ErrObjectNotFound {
				delete(endpoint.TeamAccessPolicies, policyID)
			}
		}

		err := m.endpointService.UpdateEndpoint(endpoint.ID, &endpoint)
		if err != nil {
			return err
		}
	}

	endpointGroups, err := m.endpointGroupService.EndpointGroups()
	if err != nil {
		return err
	}

	for _, endpointGroup := range endpointGroups {
		for policyID := range endpointGroup.UserAccessPolicies {
			_, err := m.userService.User(policyID)
			if err == portainer.ErrObjectNotFound {
				delete(endpointGroup.UserAccessPolicies, policyID)
			}
		}

		for policyID := range endpointGroup.TeamAccessPolicies {
			_, err := m.teamService.Team(policyID)
			if err == portainer.ErrObjectNotFound {
				delete(endpointGroup.TeamAccessPolicies, policyID)
			}
		}

		err := m.endpointGroupService.UpdateEndpointGroup(endpointGroup.ID, &endpointGroup)
		if err != nil {
			return err
		}
	}

	registries, err := m.registryService.Registries()
	if err != nil {
		return err
	}

	for _, registry := range registries {
		for policyID := range registry.UserAccessPolicies {
			_, err := m.userService.User(policyID)
			if err == portainer.ErrObjectNotFound {
				delete(registry.UserAccessPolicies, policyID)
			}
		}

		for policyID := range registry.TeamAccessPolicies {
			_, err := m.teamService.Team(policyID)
			if err == portainer.ErrObjectNotFound {
				delete(registry.TeamAccessPolicies, policyID)
			}
		}

		err := m.registryService.UpdateRegistry(registry.ID, &registry)
		if err != nil {
			return err
		}
	}

	return nil
}
