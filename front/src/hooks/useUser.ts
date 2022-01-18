import {User} from '../types';
import {useStore} from 'effector-react/effector-react.cjs';
import {$userStore} from '../store/user';

export const useUser = (): User | null => {
    const {user} = useStore($userStore);
    return user;
};
