import * as React from 'react';
import {useUserProfile} from './hooks';
import {LoadableContent} from '../../components/LoadableContent';

export type UserPageProps = {};

export const ProfilePage: React.FC<UserPageProps> = () => {
    const {userProfile, isFetch} = useUserProfile();

    return (
        <LoadableContent isLoading={isFetch}>
            <div>{userProfile && userProfile.name}</div>
        </LoadableContent>
    );
};
