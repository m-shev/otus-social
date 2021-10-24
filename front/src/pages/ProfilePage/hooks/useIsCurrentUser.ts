import {UserProfile} from '../../../types';
import {useStore} from 'effector-react';
import {$userStore} from '../../../store/user';

export const useIsCurrentUser = (userProfile: UserProfile | null): boolean => {
    const {user} = useStore($userStore);

    return user?.id === userProfile?.id;
};
