import {logoutGet} from '../../api';
import {useCallback} from 'react';
import {userLogoutEvent} from '../../store/user';
import {useHistory} from 'react-router';
import {UserProfile} from '../../types';

export interface IUseLogout {
    onLogout: () => Promise<void>;
}

export const useOnLogout = (user: UserProfile | null): IUseLogout => {
    const history = useHistory();

    const onLogout = useCallback(async () => {
        try {
            if (user) {
                await logoutGet();
                userLogoutEvent();
            }

            history.push('/login');
        } catch (e) {
            console.log(e);
        }
    }, [history, user]);

    return {
        onLogout,
    };
};
