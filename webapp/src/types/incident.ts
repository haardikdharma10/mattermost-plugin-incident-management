// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {Playbook, isPlaybook} from './playbook';

export interface Incident {
    id: string;
    name: string;
    is_active: boolean;
    commander_user_id: string;
    team_id: string;
    channel_ids: string[];
    created_at: number;
    post_id?: string;
    playbook: Playbook;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function isIncident(arg: any): arg is Incident {
    const optional = arg.post_id ? typeof arg.post_id === 'string' : true as boolean;
    return arg &&
        arg.id && typeof arg.id === 'string' &&
        arg.name && typeof arg.name === 'string' &&
        typeof arg.is_active === 'boolean' &&
        arg.commander_user_id && typeof arg.commander_user_id === 'string' &&
        arg.team_id && typeof arg.team_id === 'string' &&
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        arg.channel_ids && Array.isArray(arg.channel_ids) && arg.channel_ids.every((item: any) => typeof item === 'string') &&
        arg.created_at && typeof arg.created_at === 'number' &&
        optional &&
        arg.playbook && isPlaybook(arg.playbook);
}

export interface IncidentHeader {
    id: string;
    name: string;
    is_active: boolean;
    commander_user_id: string;
    team_id: string;
    created_at: number;
    ended_at: number;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function isIncidentHeader(arg: any): arg is IncidentHeader {
    return arg &&
        arg.id && typeof arg.id === 'string' &&
        arg.name && typeof arg.name === 'string' &&
        typeof arg.is_active === 'boolean' &&
        arg.commander_user_id && typeof arg.commander_user_id === 'string' &&
        arg.team_id && typeof arg.team_id === 'string' &&
        arg.created_at && typeof arg.created_at === 'number' &&
        arg.ended_at && typeof arg.ended_at === 'number';
}
