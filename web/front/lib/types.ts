// types.go

export type Topic = {
    id: string
    name: string
    type: TopicType
}

export type ChoiceGeneric = {
    id: string
    topic_id: string
    type: 'generic'
    order: number
    text: string
}

export type ChoiceCalendar = {
    id: string
    topic_id: string
    type: 'calendar'
    order: number
    is_all_day: boolean
    start_datetime: string
    end_datetime: string
}

export type Vote = {
    id: string
    user_name: string
    choice_id: string
}

export type TopicType = 'generic' | 'calendar'
