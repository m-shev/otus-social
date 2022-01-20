import {useLocation} from 'react-router';
import * as queryString from 'query-string';

export const useUserId = (): number => {
    const location = useLocation();
    const {id} = queryString.parse(location.search);

    return Number(id);
};
