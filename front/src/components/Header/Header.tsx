import * as React from 'react';
import styles from './Hedear.module.scss';
import {useStore} from 'effector-react';
import {$userStore} from '../../store/user';
import {Link} from 'react-router-dom';
import {useOnLogout} from './useOnLogout';

export type HeaderProps = {
    showLoginButton?: boolean;
};

export const Header: React.FC<HeaderProps> = ({showLoginButton = true}) => {
    const {user} = useStore($userStore);
    const {onLogout} = useOnLogout(user);

    return (
        <div className={styles.root}>
            <div className={styles.title}>Социальная сеть</div>
            <div className={styles.right}>
                {user && (
                    <Link className={styles.link} to={`/profile?id=${user.id}`}>
                        {user.name}
                    </Link>
                )}
                {showLoginButton && (
                    <div onClick={onLogout} className={styles.link}>
                        {user ? 'Выйти' : 'Войти'}
                    </div>
                )}
            </div>
        </div>
    );
};
