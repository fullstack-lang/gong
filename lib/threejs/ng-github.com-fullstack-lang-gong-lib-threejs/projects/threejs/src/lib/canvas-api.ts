// insertion point for imports
import { DirectionalLightAPI } from './directionallight-api'
import { AmbiantLightAPI } from './ambiantlight-api'
import { MeshAPI } from './mesh-api'
import { CameraAPI } from './camera-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CanvasAPI {

	static GONGSTRUCT_NAME = "Canvas"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	CanvasPointersEncoding: CanvasPointersEncoding = new CanvasPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class CanvasPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	DirectionalLights: number[] = []
	AmbiantLightID: NullInt64 = new NullInt64 // if pointer is null, AmbiantLight.ID = 0

	Meshs: number[] = []
	CameraID: NullInt64 = new NullInt64 // if pointer is null, Camera.ID = 0

}
