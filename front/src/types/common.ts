export type RequestState = null | 'success' | 'fail';
export enum HttpStatus {
    Ok = 200,
    BadRequest = 400,
    InternalServerError = 500,
}
