import {HttpStatus, IUseRequestState, RequestState, SetValue, UserProfile} from '../../../types';
import {useRequest} from '../../../hooks';
import {useLocation} from 'react-router';
import * as queryString from 'query-string';
import {useEffect, useState} from 'react';
import {userProfileGet} from '../../../api';
import {useStore} from 'effector-react';
import {$userStore} from '../../../store/user';

interface ILoadProfile {
    (): Promise<void>;
}

export interface IUseUserProfile extends IUseRequestState {
    userProfile: UserProfile | null;
}

const useLoadProfile = (
    id: unknown,
    setIsFetch: SetValue<boolean>,
    setRequestState: SetValue<RequestState>,
    setError: SetValue<Error | null>,
    setUserProfile: SetValue<UserProfile | null>,
): ILoadProfile => {
    return async (): Promise<void> => {
        setIsFetch(true);

        try {
            const resp = await userProfileGet(id as string);

            if (resp.status === HttpStatus.Ok) {
                setRequestState('success');
                setUserProfile(await resp.json());
            } else {
                setRequestState('fail');
                setError(new Error(await resp.text()));
            }
        } catch (e) {
            setRequestState('fail');
            setError(e as Error);
        } finally {
            setIsFetch(false);
        }
    };
};

export const useUserProfile = (): IUseUserProfile => {
    const {setIsFetch, setRequestState, setError, requestState, error, isFetch} = useRequest();
    const {user} = useStore($userStore);
    const [userProfile, setUserProfile] = useState<UserProfile | null>(null);
    const location = useLocation();
    const {id} = queryString.parse(location.search);

    const loadProfile = useLoadProfile(
        id as string,
        setIsFetch,
        setRequestState,
        setError,
        setUserProfile,
    );

    useEffect(() => {
        setRequestState(null);
        setError(null);
        setUserProfile(null);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [id]);

    useEffect(() => {
        if (user && user.id === parseInt(id as string, 10)) {
            // eslint-disable-next-line @typescript-eslint/no-unused-vars
            const {email, ...rest} = user;
            setUserProfile(rest);
        } else {
            // noinspection JSIgnoredPromiseFromCall
            loadProfile();
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return {
        error,
        isFetch,
        userProfile,
        requestState,
    };
};
