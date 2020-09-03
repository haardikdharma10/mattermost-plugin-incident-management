// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {GetStateFunc} from 'mattermost-redux/types/actions';
import {WebSocketMessage} from 'mattermost-redux/actions/websocket';
import {getCurrentTeam} from 'mattermost-redux/selectors/entities/teams';

import {Dispatch} from 'redux';

import {navigateToUrl} from 'src/browser_routing';

import {incidentCreated, receivedTeamIncidentChannels} from 'src/actions';
import {fetchIncidentChannels} from 'src/client';

import {clientId} from './selectors';
import {isIncident, Incident} from './types/incident';

export const websocketSubscribers = new Set<(incident: Incident) => void>();

export function handleReconnect(getState: GetStateFunc, dispatch: Dispatch) {
    return async (): Promise<void> => {
        const currentTeam = getCurrentTeam(getState());
        dispatch(receivedTeamIncidentChannels(await fetchIncidentChannels(currentTeam.id)));
    };
}

export function handleWebsocketIncidentUpdate() {
    return (msg: WebSocketMessage): void => {
        if (!msg.data.payload) {
            return;
        }
        const incident = JSON.parse(msg.data.payload);
        if (!isIncident(incident)) {
            return;
        }

        websocketSubscribers.forEach((fn) => fn(incident));
    };
}

export function handleWebsocketIncidentCreate(getState: GetStateFunc, dispatch: Dispatch) {
    return (msg: WebSocketMessage): void => {
        if (!msg.data.payload) {
            return;
        }
        const payload = JSON.parse(msg.data.payload);
        const incident = payload.incident;
        if (!isIncident(incident)) {
            return;
        }

        dispatch(incidentCreated(incident));

        if (payload.client_id !== clientId(getState())) {
            return;
        }

        const currentTeam = getCurrentTeam(getState());

        // Navigate to the newly created channel
        const url = `/${currentTeam.name}/channels/${incident.primary_channel_id}`;
        navigateToUrl(url);
    };
}
