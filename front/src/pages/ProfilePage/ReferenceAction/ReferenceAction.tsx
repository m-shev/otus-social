import * as React from 'react';
import {useUserReference} from '../hooks/useUserReference';
import {UserProfile} from '../../../types';
import {LoadableContent, SmallDots} from '../../../components/LoadableContent';
import styles from './ReferenceAction.module.scss';

export type ReferenceActionProps = {
    userProfile: UserProfile | null;
};

export const ReferenceAction: React.FC<ReferenceActionProps> = ({userProfile}) => {
    const {action, showAction, isFetch} = useUserReference(userProfile);
    return showAction ? (
        <div className={styles.referenceAction}>
            <LoadableContent Loader={<SmallDots />} isLoading={isFetch}>
                <button onClick={action}>Добавить в друзья</button>
            </LoadableContent>
        </div>
    ) : null;
};
