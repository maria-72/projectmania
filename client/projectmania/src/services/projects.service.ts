import { Project } from '../models/project'

class ProjectsService {
    url = process.env.REACT_APP_BASE_URL ?? ''

    getProjects = async (): Promise<Project[]> => {
        const response = await fetch(`${this.url}/api/v1/get-projects`)

        if (response.ok) {
            const jsonValue = await response.json()
            return Promise.resolve(jsonValue)
        } else {
            return Promise.reject('Error')
        }
    }
}

export const projectsService = new ProjectsService()
