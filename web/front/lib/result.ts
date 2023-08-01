export type Result<Ok, Err> = {
    Ok: () => Ok
    Err: () => Err | null
}

export function Ok<T>(value: T): Result<T, any> {
    return {
        Ok() {
            return value
        },
        Err() {
            return null
        },
    }
}

export function Err<E>(value: E): Result<any, Exclude<E, null>> {
    return {
        Ok() {
            return null
        },
        Err() {
            return value as any
        },
    }
}
