export type RequestState = null | 'success' | 'fail';
export enum HttpStatus {
    Ok = 200,
    NoContent = 204,
    BadRequest = 400,
    InternalServerError = 500,
}

export interface SetValue<V> {
    (value: V): void;
}

export interface IUseRequest {
    isFetch: boolean;
    setIsFetch: SetValue<boolean>;
    error: Error | null;
    setError: SetValue<Error | null>;
    requestState: RequestState;
    setRequestState: SetValue<RequestState>;
}

export type IUseRequestState = Pick<IUseRequest, 'isFetch' | 'requestState' | 'error'>;
