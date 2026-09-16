import {
  Injectable,
  RendererFactory2,
  RendererType2,
  Renderer2,
  makeEnvironmentProviders,
  EnvironmentProviders,
} from '@angular/core';
import { ɵDomRendererFactory2 } from '@angular/platform-browser';
import { NgtRendererFactory2, NGT_RENDERER_OPTIONS } from 'angular-three';

function isThreeComponent(type: RendererType2 | null): boolean {
  if (!type) {
    return false;
  }

  const cmpType = (type as any).type;
  if (!cmpType) {
    return false;
  }

  // 1. Explicit marker on component class
  if (cmpType.__is_three__) {
    return true;
  }

  // 2. Class name check
  const name = cmpType.name || '';
  if (name.includes('Three') || name.includes('three') || name.includes('Ngt')) {
    return true;
  }

  // 3. Selectors check on Angular Ivy ComponentDef (cmpType.ɵcmp)
  const selectors = cmpType.ɵcmp?.selectors;
  if (Array.isArray(selectors)) {
    return selectors.some((sel: any[]) =>
      sel.some(
        (tag) =>
          typeof tag === 'string' &&
          (tag.startsWith('ngt-') || tag.includes('three'))
      )
    );
  }

  return false;
}

@Injectable()
export class ScopedNgtRendererFactory2 implements RendererFactory2 {
  constructor(
    private domRendererFactory: RendererFactory2,
    private ngtRendererFactory: NgtRendererFactory2
  ) {}

  createRenderer(hostElement: any, type: RendererType2 | null): Renderer2 {
    if (isThreeComponent(type)) {
      return this.ngtRendererFactory.createRenderer(hostElement, type);
    }
    return this.domRendererFactory.createRenderer(hostElement, type);
  }

  begin(): void {
    this.domRendererFactory.begin?.();
    this.ngtRendererFactory.begin?.();
  }

  end(): void {
    this.ngtRendererFactory.end?.();
    this.domRendererFactory.end?.();
  }

  whenRenderingDone?(): Promise<any> {
    return this.domRendererFactory.whenRenderingDone?.() ?? Promise.resolve();
  }
}

export function provideScopedNgtRenderer(options = {}): EnvironmentProviders {
  return makeEnvironmentProviders([
    { provide: NGT_RENDERER_OPTIONS, useValue: options },
    {
      provide: RendererFactory2,
      useFactory: (domRendererFactory: RendererFactory2) => {
        const ngtRendererFactory = new NgtRendererFactory2(domRendererFactory);
        return new ScopedNgtRendererFactory2(
          domRendererFactory,
          ngtRendererFactory
        );
      },
      deps: [ɵDomRendererFactory2],
    },
  ]);
}
