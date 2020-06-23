// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {RouteComponentProps} from 'react-router-dom';

import Toggle from 'src/components/widgets/toggle';

import {Playbook, Checklist, ChecklistItem} from 'src/types/playbook';
import {savePlaybook} from 'src/client';
import {ChecklistDetails} from 'src/components/checklist/checklist';
import ConfirmModal from 'src/components/widgets/confirmation_modal';
import {MAX_NAME_LENGTH} from 'src/utils/constants';

import BackIcon from 'src/components/assets/icons/back_icon';

import '../playbook.scss';

export interface Props extends RouteComponentProps {
    playbook: Playbook;
    newPlaybook: boolean;
    currentTeamID: string;
    onClose: () => void;
    actions: {
        getPlaybook: (playbookId: String) => void;
    };
}

interface State{
    title: string;
    checklists: Checklist[];
    changesMade: boolean;
    confirmOpen: boolean;
    public: boolean;
}

export class PlaybookEdit extends React.PureComponent<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            title: this.props.playbook.title,
            checklists: JSON.parse(JSON.stringify(this.props.playbook.checklists)),
            changesMade: false,
            confirmOpen: false,
            public: this.props.playbook.create_public_incident,
        };
    }

    componentDidMount() {
        // If the user directly navigates to /playbooks/:playbookId, the store may not
        // contain the playbook requested. This happens only when we are not creating
        // a new playbook but the playbook in the props is still empty (it has no ID).
        if (!this.props.newPlaybook && !this.props.playbook.id) {
            this.props.actions.getPlaybook(this.props.match.params.playbookId);
        }
    }

    componentDidUpdate(prevProps) {
        if (this.props.playbook.id !== prevProps.playbook.id) {
            // eslint-disable-next-line react/no-did-update-set-state
            this.setState({
                title: this.props.playbook.title,
                checklists: JSON.parse(JSON.stringify(this.props.playbook.checklists)),
                public: this.props.playbook.create_public_incident,
            });
        }
    }

    public onSave = async () => {
        const newPlaybook: Playbook = {
            id: this.props.playbook.id,
            title: this.state.title,
            team_id: this.props.currentTeamID,
            create_public_incident: this.state.public,
            checklists: this.state.checklists,
        };

        await savePlaybook(newPlaybook);

        this.props.onClose();
    };

    public updateChecklist(newChecklist: Checklist[]) {
        this.setState({
            checklists: newChecklist,
            changesMade: true,
        });
    }

    public onAddItem = (checklistItem: ChecklistItem, checklistIndex: number): void => {
        const allChecklists = Object.assign([], this.state.checklists) as Checklist[];
        const changedChecklist = Object.assign({}, this.state.checklists[checklistIndex]);

        changedChecklist.items = [...changedChecklist.items, checklistItem];
        allChecklists[checklistIndex] = changedChecklist;

        this.updateChecklist(allChecklists);
    }

    public onDeleteItem = (checklistItemIndex: number, checklistIndex: number): void => {
        const allChecklists = Object.assign([], this.state.checklists) as Checklist[];
        const changedChecklist = Object.assign({}, allChecklists[checklistIndex]) as Checklist;

        changedChecklist.items = [
            ...changedChecklist.items.slice(0, checklistItemIndex),
            ...changedChecklist.items.slice(checklistItemIndex + 1, changedChecklist.items.length)];
        allChecklists[checklistIndex] = changedChecklist;

        this.updateChecklist(allChecklists);
    }

    public onEditItem = (checklistItemIndex: number, newTitle: string, checklistIndex: number): void => {
        const allChecklists = Object.assign([], this.state.checklists) as Checklist[];
        const changedChecklist = Object.assign({}, allChecklists[checklistIndex]) as Checklist;

        changedChecklist.items[checklistItemIndex].title = newTitle;
        allChecklists[checklistIndex] = changedChecklist;

        this.updateChecklist(allChecklists);
    }

    public onReorderItem = (checklistItemIndex: number, newIndex: number, checklistIndex: number): void => {
        const allChecklists = Object.assign([], this.state.checklists) as Checklist[];
        const changedChecklist = Object.assign({}, allChecklists[checklistIndex]) as Checklist;

        const itemToMove = changedChecklist.items[checklistItemIndex];

        // Remove from current index
        changedChecklist.items = [
            ...changedChecklist.items.slice(0, checklistItemIndex),
            ...changedChecklist.items.slice(checklistItemIndex + 1, changedChecklist.items.length)];

        // Add in new index
        changedChecklist.items = [
            ...changedChecklist.items.slice(0, newIndex),
            itemToMove,
            ...changedChecklist.items.slice(newIndex, changedChecklist.items.length + 1)];

        allChecklists[checklistIndex] = changedChecklist;

        this.updateChecklist(allChecklists);
    }

    public handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({
            title: e.target.value,
            changesMade: true,
        });
    }

    public confirmOrClose = () => {
        if (this.state.changesMade) {
            this.setState({
                confirmOpen: true,
            });
        } else {
            this.props.onClose();
        }
    }

    public confirmCancel = () => {
        this.setState({
            confirmOpen: false,
        });
    }

    public handlePublicChange = () => {
        this.setState({
            public: !this.state.public,
            changesMade: true,
        });
    }

    public render(): JSX.Element {
        const saveDisabled = this.state.title.trim() === '' || !this.state.changesMade;

        const notFoundComponent = (
            <div className='Playbook'>
                <div className='Backstage__header'>
                    <div className='title'>
                        <BackIcon
                            className='Backstage__header__back'
                            onClick={this.confirmOrClose}
                        />
                        {'Playbook Not Found'}
                    </div>
                </div>
            </div>
        );

        const editComponent = (
            <div className='Playbook'>
                <div className='Backstage__header'>
                    <div className='title'>
                        <BackIcon
                            className='Backstage__header__back'
                            onClick={this.confirmOrClose}
                        />
                        {this.props.newPlaybook ? 'New Playbook' : 'Edit Playbook'}
                    </div>
                    <div className='header-button-div'>
                        <button
                            className='btn btn-link mr-2'
                            onClick={this.confirmOrClose}
                        >
                            {'Cancel'}
                        </button>
                        <button
                            className='btn btn-primary'
                            disabled={saveDisabled}
                            onClick={this.onSave}
                        >
                            {'Save Playbook'}
                        </button>
                    </div>
                </div>
                <div className='playbook-fields'>
                    <input
                        autoFocus={true}
                        id={'playbook-name'}
                        className='form-control input-name'
                        type='text'
                        placeholder='Playbook Name'
                        value={this.state.title}
                        maxLength={MAX_NAME_LENGTH}
                        onChange={this.handleTitleChange}
                    />
                    <div className='public-item'>
                        <div
                            className='checkbox-container'
                        >
                            <Toggle
                                toggled={this.state.public}
                                onToggle={this.handlePublicChange}
                            />
                            <label>
                                {'Create Public Incident'}
                            </label>
                        </div>
                    </div>
                    <div className='checklist-container'>
                        {this.state.checklists?.map((checklist: Checklist, checklistIndex: number) => (
                            <ChecklistDetails
                                checklist={checklist}
                                enableEdit={true}
                                key={checklist.title + checklistIndex}

                                addItem={(checklistItem: ChecklistItem) => {
                                    this.onAddItem(checklistItem, checklistIndex);
                                }}
                                removeItem={(chceklistItemIndex: number) => {
                                    this.onDeleteItem(chceklistItemIndex, checklistIndex);
                                }}
                                editItem={(checklistItemIndex: number, newTitle: string) => {
                                    this.onEditItem(checklistItemIndex, newTitle, checklistIndex);
                                }}
                                reorderItems={(checklistItemIndex: number, newPosition: number) => {
                                    this.onReorderItem(checklistItemIndex, newPosition, checklistIndex);
                                }}
                            />
                        ))}
                    </div>
                </div>
                <ConfirmModal
                    show={this.state.confirmOpen}
                    title={'Confirm discard'}
                    message={'Are you sure you want to discard your changes?'}
                    confirmButtonText={'Discard Changes'}
                    onConfirm={this.props.onClose}
                    onCancel={this.confirmCancel}
                />
            </div>
        );

        const isPlaybookDefined = this.props.newPlaybook || this.props.playbook.id;

        return isPlaybookDefined ? editComponent : notFoundComponent;
    }
}