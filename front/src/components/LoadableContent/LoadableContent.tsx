import * as React from 'react';
import {ReactNode} from 'react';
import {DefaultLoader} from './DefaultLoader';

export type Props = {
    children: ReactNode;
    isLoading: boolean;
    Loader?: JSX.Element;
};

export function LoadableContent({
    children,
    isLoading,
    Loader = <DefaultLoader />,
}: Props): JSX.Element {
    return isLoading ? <>{Loader}</> : <>{children}</>;
}
