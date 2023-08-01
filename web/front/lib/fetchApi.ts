import {
    ReesponseGetVote,
    RequestCreateTopicCalendar,
    RequestCreateTopicGeneric,
    RequestCreateVote,
    ResponseCreateTopicCalendar,
    ResponseCreateTopicGenreic,
    ResponseCreateVote,
    ResponseGetTopicCalendar,
    ResponseGetTopicGeneric,
} from './apiTypes'
import { Err, Ok, Result } from './result'

type endpoints = {
    '/api/topic/get/:id': [
        { id: string },
        {},
        ResponseGetTopicGeneric | ResponseGetTopicCalendar,
        'GET',
    ]
    '/api/vote/get/:id': [{ id: string }, {}, ReesponseGetVote, 'GET']
    '/api/vote/new': [{}, RequestCreateVote, ResponseCreateVote, 'POST']
    '/api/topic/new/generic': [
        {},
        RequestCreateTopicGeneric,
        ResponseCreateTopicGenreic,
        'POST',
    ]
    '/api/topic/new/calendar': [
        {},
        RequestCreateTopicCalendar,
        ResponseCreateTopicCalendar,
        'POST',
    ]
}

function resolveParam(
    path: string,
    params: { [key: string]: string },
): Result<string, string> {
    const paths = path.split('/')
    const replacedPaths: string[] = []

    for (const path of paths) {
        if (!path.startsWith(':')) {
            replacedPaths.push(path)
            continue
        }
        const key = path.slice(1)
        const value = params[key]
        if (!value) {
            return Err('invalid parameter')
        }
        replacedPaths.push(value)
    }

    return Ok(replacedPaths.join('/'))
}

export async function fetchApi<Path extends keyof endpoints>(
    path: Path,
    method: endpoints[Path][3],
    param: endpoints[Path][0],
    body: endpoints[Path][1],
): Promise<Result<endpoints[Path][2], string>> {
    const resolvedPath = resolveParam(path, param)
    if (resolvedPath.Err()) {
        return Err(resolvedPath.Err())
    }

    try {
        const json = await fetch(resolvedPath.Ok(), {
            method,
            body: JSON.stringify(body),
        }).then((res) => res.json())

        if (json.error) {
            return Err(json.message)
        }
        return Ok(json.body)
    } catch (err) {
        return Err(err.toString())
    }
}
