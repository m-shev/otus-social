import {useStore} from 'effector-react';
import {$userStore} from '../../../store/user';
import {HttpStatus, IUseRequestState, UserFriend, UserProfile} from '../../../types';
import {useCallback, useMemo} from 'react';
import {useRequest} from '../../../hooks';
import {addFriendPost, removeFriendDelete} from '../../../api';
import {ILoadProfile} from './useUserProfile';

interface IAction {
    (): Promise<void>;
}

const ADD_TO_FRIEND = 'Добавить в друзья';
const REMOVE_FROM_FRIEND = 'Удалить из друзей';

export interface IUseUserReference extends IUseRequestState {
    showAction: boolean;
    action: IAction;
    actionName: string;
}

const useIsFriend = (userId: number | undefined, friends: UserFriend[]): boolean => {
    return useMemo(() => {
        return friends.some((friend) => {
            return friend.id === userId;
        });
    }, [userId, friends]);
};

export const useUserReference = (
    userProfile: UserProfile | null,
    loadProfile: ILoadProfile,
): IUseUserReference => {
    const {setIsFetch, setError, setRequestState, isFetch, requestState, error} = useRequest();
    const {user} = useStore($userStore);

    let showAction = false;

    if (userProfile && user && user.id !== userProfile.id) {
        showAction = true;
    }

    const addFriendAction = useCallback(async () => {
        try {
            setIsFetch(true);
            const resp = await addFriendPost({
                userId: user?.id || 0,
                friendId: userProfile?.id || 0,
            });

            if (resp.status === HttpStatus.Ok) {
                setRequestState('success');
                await loadProfile();
            } else {
                throw new Error(await resp.text());
            }
        } catch (e) {
            setRequestState('fail');
            e instanceof Error ? setError(e) : setError(new Error(`Неизвестная ошибка`));
        } finally {
            setIsFetch(false);
        }
    }, [loadProfile, setError, setIsFetch, setRequestState, user?.id, userProfile?.id]);

    const removeFromFriend = useCallback(async () => {
        try {
            const resp = await removeFriendDelete({
                userId: user?.id || 0,
                friendId: userProfile?.id || 0,
            });

            if (resp.status === HttpStatus.Ok) {
                setRequestState('success');
                await loadProfile();
            } else {
                throw new Error(await resp.text());
            }
        } catch (e) {
            setRequestState('fail');
            e instanceof Error ? setError(e) : setError(new Error(`Неизвестная ошибка`));
        } finally {
            setIsFetch(false);
        }
    }, [loadProfile, setError, setIsFetch, setRequestState, user?.id, userProfile?.id]);

    const isFriend = useIsFriend(user?.id, userProfile?.friends || []);

    return {
        showAction,
        action: isFriend ? removeFromFriend : addFriendAction,
        actionName: isFriend ? REMOVE_FROM_FRIEND : ADD_TO_FRIEND,
        isFetch,
        error,
        requestState,
    };
};
