import {useStore} from 'effector-react';
import {$userStore} from '../../../store/user';
import {HttpStatus, IUseRequestState, UserProfile} from '../../../types';
import {useCallback} from 'react';
import {useRequest} from '../../../hooks';
import {addFriendPost} from '../../../api';

export interface IUseUserReference extends IUseRequestState {
    showAction: boolean;
    action: () => Promise<void>;
}

export const useUserReference = (userProfile: UserProfile | null): IUseUserReference => {
    const {setIsFetch, setError, setRequestState, isFetch, requestState, error} = useRequest();
    const {user} = useStore($userStore);

    let showAction = false;

    if (userProfile && user?.id !== userProfile.id) {
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
            } else {
                throw new Error(await resp.text());
            }
        } catch (e) {
            setRequestState('fail');
            e instanceof Error ? setError(e) : setError(new Error(`Неизвестная ошибка`));
        } finally {
            setIsFetch(false);
        }
    }, [setError, setIsFetch, setRequestState, user?.id, userProfile?.id]);

    return {
        showAction,
        action: addFriendAction,
        isFetch,
        error,
        requestState,
    };
};
