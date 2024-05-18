import {
    ArrowRightOnRectangleIcon,
    ShieldCheckIcon
} from "@heroicons/react/24/solid"
import { useMutateAuth } from "../hooks/useMutateAuth"
import { useQueryTasks } from "../hooks/useQueryTasks"
import { Task } from "../types"
import { TaskItem } from "./TaskItem"
import { useQueryClient } from "@tanstack/react-query"
import useStore from "../store"
import { useMutateTasks } from "../hooks/useMutateTask"
import { FormEvent, useEffect } from "react"
import { setInterval } from "timers/promises"

export const ToDo = () => {
    const { logoutMutating } = useMutateAuth()
    const queryClient = useQueryClient()
    const { editedTask } = useStore()
    const { data, isLoading, refetch } = useQueryTasks()
    const updateTask = useStore((state) => state.updateEditedTask)
    const { createTaskMutation, updateTaskMutation } = useMutateTasks()

    const logout = async () => {
        await logoutMutating.mutateAsync()
        queryClient.removeQueries(['tasks'])
    }

    const submitTaskHandler = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (editedTask.id === 0) {
            createTaskMutation.mutate({ title: editedTask.title })
        } else {
            updateTaskMutation.mutate(editedTask)
        }
    }

    const submitTaskRefetchHandler = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        refetch()
    }

    return (
        <div className="flex justify-center items-center flex-col min-h-screen text-grey-600 font-mono">
            <div className="flex items-center">
                <span>log out</span>
                <ArrowRightOnRectangleIcon className="h-6 w-6 my-6 text-blue-500 cursor-pointer" onClick={logout} />
            </div>
            <div className="flex items-center">
                <ShieldCheckIcon className="w-8 h-8 mr-2 text-blue-500 cursor-pointer" />
                <span className="font-bold text-2xl">Task Manager</span>
                <form onSubmit={submitTaskRefetchHandler}>
                    <button
                        className="py-2 px-2 ml-10 mt-4 bg-blue-600 hover:bg-green-700 text-sm rounded-md text-white disabled:opacity-40"
                        type="submit">最新化
                    </button>
                </form>
            </div>
            <form onSubmit={submitTaskHandler}>
                <input
                    type="text"
                    className="w-full px-4 mt-4 py-2 border rounded-md"
                    placeholder="New task ?"
                    value={editedTask.title || ''}
                    onChange={(e) => updateTask({ id: editedTask.id, title: e.target.value })} />
                <button
                    className="w-full py-2 mt-4 bg-blue-600 hover:bg-green-700 rounded-md text-white disabled:opacity-40"
                    type="submit"
                    disabled={!editedTask.title}>
                    {editedTask.id === 0 ? 'Create Task' : 'Update Task'}</button>
            </form>
            {isLoading ? (<p></p>) : (
                <ul className="my-5">
                    {data?.map((task: Task) => (
                        <TaskItem key={task.id} id={task.id} title={task.title} />
                    ))}
                </ul>
            )}
        </div>
    )
}

export default ToDo
