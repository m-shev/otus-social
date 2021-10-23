import {useStore} from 'effector-react';
import {$userStore} from '../store/user';
import {useHistory} from 'react-router';
import {useEffect} from 'react';

export const useRedirectToProfile = (): void => {
    const {user} = useStore($userStore);
    const history = useHistory();

    useEffect(() => {
        if (user != null) {
            history.push(`profile?id=${user.id}`);
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [user]);
};
