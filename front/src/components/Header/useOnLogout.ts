import {logout} from '../../api';
import {useCallback} from 'react';
import {userLogoutEvent} from '../../store/user';
import {useHistory} from 'react-router';
import {User} from '../../types';

export interface IUseLogout {
    onLogout: () => Promise<void>;
}

export const useOnLogout = (user: User | null): IUseLogout => {
    const history = useHistory();

    const onLogout = useCallback(async () => {
        try {
            if (user) {
                await logout();
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
