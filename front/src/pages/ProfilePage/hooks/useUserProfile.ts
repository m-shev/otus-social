import {HttpStatus, IUseRequestState, RequestState, SetValue, UserProfile} from '../../../types';
import {useRequest} from '../../../hooks';
import {useLocation} from 'react-router';
import * as queryString from 'query-string';
import {useEffect, useState} from 'react';
import {friendListGet, profileGet} from '../../../api';

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
        setRequestState(null);

        try {
            const resp = await profileGet(id as string);
            let profile: UserProfile | null = null;

            if (resp.status === HttpStatus.Ok) {
                profile = await resp.json();
            } else {
                setRequestState('fail');
                throw new Error(await resp.text());
            }

            const friendResp = await friendListGet(id as string);

            if (profile && friendResp.status === HttpStatus.Ok) {
                setRequestState('success');
                profile.friends = (await friendResp.json()).friendList;
                setUserProfile(profile);
            } else {
                throw new Error(await friendResp.text());
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
        // noinspection JSIgnoredPromiseFromCall
        loadProfile();
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [location]);

    return {
        error,
        isFetch,
        userProfile,
        requestState,
    };
};
