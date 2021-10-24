import {UserProfile} from '../../../types';

export const useFriendText = (userProfile: UserProfile | null, isCurrentUser: boolean): string => {
    if (!userProfile) {
        return '';
    }

    const hasFriend = !!userProfile.friends.length;

    if (!hasFriend && isCurrentUser) {
        return 'У вас пока нет друзей';
    }

    if (!hasFriend && !isCurrentUser) {
        return `У ${userProfile.name} пока нет друзей`;
    }

    return 'Друзья';
};
