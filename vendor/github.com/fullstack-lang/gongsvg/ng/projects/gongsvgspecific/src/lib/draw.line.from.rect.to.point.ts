
import * as gongsvg from 'gongsvg'
import { createPoint } from './link/draw.segments';


// Function to calculate the center of a rectangle
function getRectCenter(rect: gongsvg.RectDB): gongsvg.PointDB {
    return createPoint(
        rect.X + rect.Width / 2,
        rect.Y + rect.Height / 2,
    )
}

// Function to draw a line from the border of the rectangle to the point B
export function drawLineFromRectToB(rect: gongsvg.RectDB, B: gongsvg.PointDB): gongsvg.PointDB {
    const center = getRectCenter(rect);

    // Calculate vector from center to point B
    const vectorX = B.X - center.X;
    const vectorY = B.Y - center.Y;

    // Normalize the vector
    const magnitude = Math.sqrt(vectorX * vectorX + vectorY * vectorY);
    const unitVectorX = vectorX / magnitude;
    const unitVectorY = vectorY / magnitude;

    // Calculate the intersection point of the line with the rectangle border
    // We do this by checking the intersection with all four sides and taking the one that falls within the rectangle bounds
    let borderX = center.X;
    let borderY = center.Y;

    // Top border
    if (unitVectorY < 0) {
        const t = (rect.Y - center.Y) / unitVectorY;
        const potentialX = center.X + t * unitVectorX;
        if (potentialX >= rect.X && potentialX <= rect.X + rect.Width) {
            borderX = potentialX;
            borderY = rect.Y;
        }
    }

    // Bottom border
    if (unitVectorY > 0) {
        const t = (rect.Y + rect.Height - center.Y) / unitVectorY;
        const potentialX = center.X + t * unitVectorX;
        if (potentialX >= rect.X && potentialX <= rect.X + rect.Width) {
            borderX = potentialX;
            borderY = rect.Y + rect.Height;
        }
    }

    // Left border
    if (unitVectorX < 0) {
        const t = (rect.X - center.X) / unitVectorX;
        const potentialY = center.Y + t * unitVectorY;
        if (potentialY >= rect.Y && potentialY <= rect.Y + rect.Height) {
            borderX = rect.X;
            borderY = potentialY;
        }
    }

    // Right border
    if (unitVectorX > 0) {
        const t = (rect.X + rect.Width - center.X) / unitVectorX;
        const potentialY = center.Y + t * unitVectorY;
        if (potentialY >= rect.Y && potentialY <= rect.Y + rect.Height) {
            borderX = rect.X + rect.Width;
            borderY = potentialY;
        }
    }

    return createPoint(borderX, borderY)
}
