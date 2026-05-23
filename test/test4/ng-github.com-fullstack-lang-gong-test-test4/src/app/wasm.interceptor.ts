import { HttpInterceptorFn, HttpResponse } from '@angular/common/http';
import { of } from 'rxjs';

export const wasmInterceptor: HttpInterceptorFn = (req, next) => {
  // Only intercept API calls if the WASM backend is loaded
  if (req.url.includes('/api/') && (window as any).wasmFetch) {
    
    const reqData = JSON.stringify({
      method: req.method,
      url: req.url,
      body: req.body ? JSON.stringify(req.body) : ''
    });

    // Pass the request to Go via the JS bridge
    const goResponseString = (window as any).wasmFetch(reqData);
    const goResponse = JSON.parse(goResponseString);

    // Reconstruct the Angular HttpResponse
    const response = new HttpResponse({
      body: goResponse.body ? JSON.parse(goResponse.body) : null,
      status: goResponse.status,
      url: req.url
    });

    return of(response);
  }
  
  // If not an API call or WASM isn't loaded, continue normally
  return next(req);
};