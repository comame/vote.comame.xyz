import { ChoiceCalendar, ChoiceGeneric, Topic, Vote } from './types'

/// httpType.go

export type TopicGenericWithChoices = {
    topic: Topic
    choices: ChoiceGeneric[]
}

export type TopicCalendarWithChoices = {
    topic: Topic
    choices: ChoiceCalendar[]
}

export type ResponseGetTopicGeneric = TopicGenericWithChoices

export type ResponseGetTopicCalendar = TopicCalendarWithChoices

export type RequestCreateTopicGeneric = TopicGenericWithChoices

export type ResponseCreateTopicGenreic = response<TopicGenericWithChoices>

export type RequestCreateTopicCalendar = TopicCalendarWithChoices

export type ResponseCreateTopicCalendar = response<TopicCalendarWithChoices>

export type RequestCreateVote = Vote

export type ResponseCreateVote = response<{}>

export type ReesponseGetVote = response<Vote[]>

type response<T> =
    | {
          error: true
          message: string
      }
    | {
          error: false
          body: T
      }
