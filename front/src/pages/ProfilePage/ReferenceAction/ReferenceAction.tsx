import * as React from 'react';
import {useUserReference} from '../hooks/useUserReference';
import {UserProfile} from '../../../types';
import {LoadableContent, SmallDots} from '../../../components/LoadableContent';
import styles from './ReferenceAction.module.scss';
import {ILoadProfile} from '../hooks';

export type ReferenceActionProps = {
    userProfile: UserProfile | null;
    loadProfile: ILoadProfile;
};

export const ReferenceAction: React.FC<ReferenceActionProps> = ({userProfile, loadProfile}) => {
    const {action, showAction, isFetch, actionName} = useUserReference(userProfile, loadProfile);

    return showAction ? (
        <div className={styles.referenceAction}>
            <LoadableContent Loader={<SmallDots />} isLoading={isFetch}>
                <button onClick={action}>{actionName}</button>
            </LoadableContent>
        </div>
    ) : null;
};
